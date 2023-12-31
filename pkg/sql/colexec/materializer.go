// Copyright 2018 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package colexec

import (
	"context"
	"fmt"

	"github.com/cockroachdb/cockroach/pkg/col/coldata"
	"github.com/cockroachdb/cockroach/pkg/sql/colexec/execerror"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfra"
	"github.com/cockroachdb/cockroach/pkg/sql/execinfrapb"
	"github.com/cockroachdb/cockroach/pkg/sql/sqlbase"
	"github.com/cockroachdb/cockroach/pkg/sql/types"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

// Materializer converts an Operator input into a execinfra.RowSource.
type Materializer struct {
	execinfra.ProcessorBase
	NonExplainable

	input Operator

	da sqlbase.DatumAlloc

	// runtime fields --

	// curIdx represents the current index into the column batch: the next row the
	// Materializer will emit.
	curIdx int
	// batch is the current Batch the Materializer is processing.
	batch coldata.Batch

	// row is the memory used for the output row.
	row sqlbase.EncDatumRow

	// outputRow stores the returned results of next() to be passed through an
	// adapter.
	outputRow sqlbase.EncDatumRow

	// cancelFlow will return a function to cancel the context of the flow. It is
	// a function in order to be lazily evaluated, since the context cancellation
	// function is only available when Starting. This function differs from
	// ctxCancel in that it will cancel all components of the Materializer's flow,
	// including those started asynchronously.
	cancelFlow func() context.CancelFunc

	// closers is a slice of IdempotentClosers that should be Closed on
	// termination.
	closers []IdempotentCloser
}

const materializerProcName = "materializer"

// NewMaterializer creates a new Materializer processor which processes the
// columnar data coming from input to return it as rows.
// Arguments:
// - typs is the output types scheme.
// - metadataSourcesQueue are all of the metadata sources that are planned on
// the same node as the Materializer and that need to be drained.
// - outputStatsToTrace (when tracing is enabled) finishes the stats.
// - cancelFlow should return the context cancellation function that cancels
// the context of the flow (i.e. it is Flow.ctxCancel). It should only be
// non-nil in case of a root Materializer (i.e. not when we're wrapping a row
// source).
func NewMaterializer(
	flowCtx *execinfra.FlowCtx,
	processorID int32,
	input Operator,
	typs []types.T,
	post *execinfrapb.PostProcessSpec,
	output execinfra.RowReceiver,
	metadataSourcesQueue []execinfrapb.MetadataSource,
	toClose []IdempotentCloser,
	outputStatsToTrace func(),
	cancelFlow func() context.CancelFunc,
) (*Materializer, error) {
	m := &Materializer{
		input:   input,
		row:     make(sqlbase.EncDatumRow, len(typs)),
		closers: toClose,
	}

	if err := m.ProcessorBase.Init(
		m,
		post,
		typs,
		flowCtx,
		processorID,
		output,
		nil, /* memMonitor */
		execinfra.ProcStateOpts{
			TrailingMetaCallback: func(ctx context.Context) []execinfrapb.ProducerMetadata {
				var trailingMeta []execinfrapb.ProducerMetadata
				for _, src := range metadataSourcesQueue {
					trailingMeta = append(trailingMeta, src.DrainMeta(ctx)...)
				}
				m.InternalClose()
				return trailingMeta
			},
		},
	); err != nil {
		return nil, err
	}
	m.FinishTrace = outputStatsToTrace
	m.cancelFlow = cancelFlow
	return m, nil
}

var _ execinfra.OpNode = &Materializer{}

// ChildCount is part of the exec.OpNode interface.
func (m *Materializer) ChildCount(verbose bool) int {
	return 1
}

// Child is part of the exec.OpNode interface.
func (m *Materializer) Child(nth int, verbose bool) execinfra.OpNode {
	if nth == 0 {
		return m.input
	}
	execerror.VectorizedInternalPanic(fmt.Sprintf("invalid index %d", nth))
	// This code is unreachable, but the compiler cannot infer that.
	return nil
}

// Start is part of the execinfra.RowSource interface.
func (m *Materializer) Start(ctx context.Context) context.Context {
	ctx = m.ProcessorBase.StartInternal(ctx, materializerProcName)
	// We can encounter an expected error during Init (e.g. an operator
	// attempts to allocate a batch, but the memory budget limit has been
	// reached), so we need to wrap it with a catcher.
	if err := execerror.CatchVectorizedRuntimeError(m.input.Init); err != nil {
		m.MoveToDraining(err)
	}
	return ctx
}

// nextAdapter calls next() and saves the returned results in m. For internal
// use only. The purpose of having this function is to not create an anonymous
// function on every call to Next().
func (m *Materializer) nextAdapter() {
	m.outputRow = m.next()
}

// next is the logic of Next() extracted in a separate method to be used by an
// adapter to be able to wrap the latter with a catcher. nil is returned when
// a zero-length batch is encountered.
func (m *Materializer) next() sqlbase.EncDatumRow {
	if m.batch == nil || m.curIdx >= m.batch.Length() {
		// Get a fresh batch.
		m.batch = m.input.Next(m.Ctx)

		if m.batch.Length() == 0 {
			return nil
		}
		m.curIdx = 0
	}
	sel := m.batch.Selection()

	rowIdx := m.curIdx
	if sel != nil {
		rowIdx = sel[m.curIdx]
	}
	m.curIdx++

	typs := m.OutputTypes()
	for colIdx := 0; colIdx < len(typs); colIdx++ {
		col := m.batch.ColVec(colIdx)
		m.row[colIdx].Datum = PhysicalTypeColElemToDatum(col, rowIdx, &m.da, &typs[colIdx])
	}
	return m.ProcessRowHelper(m.row)
}

// Next is part of the execinfra.RowSource interface.
func (m *Materializer) Next() (sqlbase.EncDatumRow, *execinfrapb.ProducerMetadata) {
	for m.State == execinfra.StateRunning {
		if err := execerror.CatchVectorizedRuntimeError(m.nextAdapter); err != nil {
			m.MoveToDraining(err)
			continue
		}
		if m.outputRow == nil {
			// Zero-length batch was encountered, move to draining.
			m.MoveToDraining(nil /* err */)
			continue
		}
		return m.outputRow, nil
	}
	// Forward any metadata.
	return nil, m.DrainHelper()
}

// InternalClose helps implement the execinfra.RowSource interface.
func (m *Materializer) InternalClose() bool {
	if m.ProcessorBase.InternalClose() {
		if m.cancelFlow != nil {
			m.cancelFlow()()
		}
		for _, closer := range m.closers {
			if err := closer.IdempotentClose(m.Ctx); err != nil {
				if log.V(1) {
					log.Infof(m.Ctx, "error closing Closer: %v", err)
				}
			}
		}
		return true
	}
	return false
}

// ConsumerDone is part of the execinfra.RowSource interface.
func (m *Materializer) ConsumerDone() {
	// Materializer will move into 'draining' state, and after all the metadata
	// has been drained - as part of TrailingMetaCallback - InternalClose() will
	// be called which will cancel the flow.
	m.MoveToDraining(nil /* err */)
}

// ConsumerClosed is part of the execinfra.RowSource interface.
func (m *Materializer) ConsumerClosed() {
	m.InternalClose()
}
