// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/enginepb/rocksdb.proto

package enginepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// SSTUserProperties contains the user-added properties of a single sstable.
type SSTUserProperties struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// ts_min is the minimum mvcc timestamp present in this sstable.
	TsMin *hlc.Timestamp `protobuf:"bytes,2,opt,name=ts_min,json=tsMin,proto3" json:"ts_min,omitempty"`
	// ts_max is the maximum mvcc timestamp present in this sstable.
	TsMax *hlc.Timestamp `protobuf:"bytes,3,opt,name=ts_max,json=tsMax,proto3" json:"ts_max,omitempty"`
}

func (m *SSTUserProperties) Reset()         { *m = SSTUserProperties{} }
func (m *SSTUserProperties) String() string { return proto.CompactTextString(m) }
func (*SSTUserProperties) ProtoMessage()    {}
func (*SSTUserProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocksdb_e0dede928993ad0c, []int{0}
}
func (m *SSTUserProperties) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SSTUserProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SSTUserProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSTUserProperties.Merge(dst, src)
}
func (m *SSTUserProperties) XXX_Size() int {
	return m.Size()
}
func (m *SSTUserProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_SSTUserProperties.DiscardUnknown(m)
}

var xxx_messageInfo_SSTUserProperties proto.InternalMessageInfo

// SSTUserPropertiesCollection contains the user-added properties of every
// sstable in a RocksDB instance.
type SSTUserPropertiesCollection struct {
	Sst   []SSTUserProperties `protobuf:"bytes,1,rep,name=sst,proto3" json:"sst"`
	Error string              `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (m *SSTUserPropertiesCollection) Reset()         { *m = SSTUserPropertiesCollection{} }
func (m *SSTUserPropertiesCollection) String() string { return proto.CompactTextString(m) }
func (*SSTUserPropertiesCollection) ProtoMessage()    {}
func (*SSTUserPropertiesCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocksdb_e0dede928993ad0c, []int{1}
}
func (m *SSTUserPropertiesCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SSTUserPropertiesCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SSTUserPropertiesCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSTUserPropertiesCollection.Merge(dst, src)
}
func (m *SSTUserPropertiesCollection) XXX_Size() int {
	return m.Size()
}
func (m *SSTUserPropertiesCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_SSTUserPropertiesCollection.DiscardUnknown(m)
}

var xxx_messageInfo_SSTUserPropertiesCollection proto.InternalMessageInfo

// HistogramData holds the relevant metrics returned by a call to
// `rocksdb::Statistics::histogramData()`.
type HistogramData struct {
	Mean  float64 `protobuf:"fixed64,1,opt,name=mean,proto3" json:"mean,omitempty"`
	P50   float64 `protobuf:"fixed64,2,opt,name=p50,proto3" json:"p50,omitempty"`
	P95   float64 `protobuf:"fixed64,3,opt,name=p95,proto3" json:"p95,omitempty"`
	P99   float64 `protobuf:"fixed64,4,opt,name=p99,proto3" json:"p99,omitempty"`
	Max   float64 `protobuf:"fixed64,5,opt,name=max,proto3" json:"max,omitempty"`
	Count uint64  `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	Sum   uint64  `protobuf:"varint,7,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (m *HistogramData) Reset()         { *m = HistogramData{} }
func (m *HistogramData) String() string { return proto.CompactTextString(m) }
func (*HistogramData) ProtoMessage()    {}
func (*HistogramData) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocksdb_e0dede928993ad0c, []int{2}
}
func (m *HistogramData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HistogramData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *HistogramData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistogramData.Merge(dst, src)
}
func (m *HistogramData) XXX_Size() int {
	return m.Size()
}
func (m *HistogramData) XXX_DiscardUnknown() {
	xxx_messageInfo_HistogramData.DiscardUnknown(m)
}

var xxx_messageInfo_HistogramData proto.InternalMessageInfo

// TickersAndHistograms holds maps from ticker/histogram name to its value for
// all stats measured by a `rocksdb::Statistics` object.
type TickersAndHistograms struct {
	Tickers    map[string]uint64        `protobuf:"bytes,1,rep,name=tickers,proto3" json:"tickers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Histograms map[string]HistogramData `protobuf:"bytes,2,rep,name=histograms,proto3" json:"histograms" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *TickersAndHistograms) Reset()         { *m = TickersAndHistograms{} }
func (m *TickersAndHistograms) String() string { return proto.CompactTextString(m) }
func (*TickersAndHistograms) ProtoMessage()    {}
func (*TickersAndHistograms) Descriptor() ([]byte, []int) {
	return fileDescriptor_rocksdb_e0dede928993ad0c, []int{3}
}
func (m *TickersAndHistograms) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickersAndHistograms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *TickersAndHistograms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickersAndHistograms.Merge(dst, src)
}
func (m *TickersAndHistograms) XXX_Size() int {
	return m.Size()
}
func (m *TickersAndHistograms) XXX_DiscardUnknown() {
	xxx_messageInfo_TickersAndHistograms.DiscardUnknown(m)
}

var xxx_messageInfo_TickersAndHistograms proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SSTUserProperties)(nil), "cockroach.storage.enginepb.SSTUserProperties")
	proto.RegisterType((*SSTUserPropertiesCollection)(nil), "cockroach.storage.enginepb.SSTUserPropertiesCollection")
	proto.RegisterType((*HistogramData)(nil), "cockroach.storage.enginepb.HistogramData")
	proto.RegisterType((*TickersAndHistograms)(nil), "cockroach.storage.enginepb.TickersAndHistograms")
	proto.RegisterMapType((map[string]HistogramData)(nil), "cockroach.storage.enginepb.TickersAndHistograms.HistogramsEntry")
	proto.RegisterMapType((map[string]uint64)(nil), "cockroach.storage.enginepb.TickersAndHistograms.TickersEntry")
}
func (m *SSTUserProperties) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSTUserProperties) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Path) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(len(m.Path)))
		i += copy(dAtA[i:], m.Path)
	}
	if m.TsMin != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(m.TsMin.Size()))
		n1, err := m.TsMin.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.TsMax != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(m.TsMax.Size()))
		n2, err := m.TsMax.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *SSTUserPropertiesCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SSTUserPropertiesCollection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Sst) > 0 {
		for _, msg := range m.Sst {
			dAtA[i] = 0xa
			i++
			i = encodeVarintRocksdb(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Error) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(len(m.Error)))
		i += copy(dAtA[i:], m.Error)
	}
	return i, nil
}

func (m *HistogramData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HistogramData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Mean != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Mean))))
		i += 8
	}
	if m.P50 != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.P50))))
		i += 8
	}
	if m.P95 != 0 {
		dAtA[i] = 0x19
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.P95))))
		i += 8
	}
	if m.P99 != 0 {
		dAtA[i] = 0x21
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.P99))))
		i += 8
	}
	if m.Max != 0 {
		dAtA[i] = 0x29
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Max))))
		i += 8
	}
	if m.Count != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(m.Count))
	}
	if m.Sum != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintRocksdb(dAtA, i, uint64(m.Sum))
	}
	return i, nil
}

func (m *TickersAndHistograms) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickersAndHistograms) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tickers) > 0 {
		keysForTickers := make([]string, 0, len(m.Tickers))
		for k := range m.Tickers {
			keysForTickers = append(keysForTickers, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForTickers)
		for _, k := range keysForTickers {
			dAtA[i] = 0xa
			i++
			v := m.Tickers[string(k)]
			mapSize := 1 + len(k) + sovRocksdb(uint64(len(k))) + 1 + sovRocksdb(uint64(v))
			i = encodeVarintRocksdb(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintRocksdb(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x10
			i++
			i = encodeVarintRocksdb(dAtA, i, uint64(v))
		}
	}
	if len(m.Histograms) > 0 {
		keysForHistograms := make([]string, 0, len(m.Histograms))
		for k := range m.Histograms {
			keysForHistograms = append(keysForHistograms, string(k))
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForHistograms)
		for _, k := range keysForHistograms {
			dAtA[i] = 0x12
			i++
			v := m.Histograms[string(k)]
			msgSize := 0
			if (&v) != nil {
				msgSize = (&v).Size()
				msgSize += 1 + sovRocksdb(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovRocksdb(uint64(len(k))) + msgSize
			i = encodeVarintRocksdb(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintRocksdb(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintRocksdb(dAtA, i, uint64((&v).Size()))
			n3, err := (&v).MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n3
		}
	}
	return i, nil
}

func encodeVarintRocksdb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SSTUserProperties) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovRocksdb(uint64(l))
	}
	if m.TsMin != nil {
		l = m.TsMin.Size()
		n += 1 + l + sovRocksdb(uint64(l))
	}
	if m.TsMax != nil {
		l = m.TsMax.Size()
		n += 1 + l + sovRocksdb(uint64(l))
	}
	return n
}

func (m *SSTUserPropertiesCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Sst) > 0 {
		for _, e := range m.Sst {
			l = e.Size()
			n += 1 + l + sovRocksdb(uint64(l))
		}
	}
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovRocksdb(uint64(l))
	}
	return n
}

func (m *HistogramData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mean != 0 {
		n += 9
	}
	if m.P50 != 0 {
		n += 9
	}
	if m.P95 != 0 {
		n += 9
	}
	if m.P99 != 0 {
		n += 9
	}
	if m.Max != 0 {
		n += 9
	}
	if m.Count != 0 {
		n += 1 + sovRocksdb(uint64(m.Count))
	}
	if m.Sum != 0 {
		n += 1 + sovRocksdb(uint64(m.Sum))
	}
	return n
}

func (m *TickersAndHistograms) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Tickers) > 0 {
		for k, v := range m.Tickers {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovRocksdb(uint64(len(k))) + 1 + sovRocksdb(uint64(v))
			n += mapEntrySize + 1 + sovRocksdb(uint64(mapEntrySize))
		}
	}
	if len(m.Histograms) > 0 {
		for k, v := range m.Histograms {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovRocksdb(uint64(len(k))) + 1 + l + sovRocksdb(uint64(l))
			n += mapEntrySize + 1 + sovRocksdb(uint64(mapEntrySize))
		}
	}
	return n
}

func sovRocksdb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRocksdb(x uint64) (n int) {
	return sovRocksdb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SSTUserProperties) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRocksdb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SSTUserProperties: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSTUserProperties: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TsMin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TsMin == nil {
				m.TsMin = &hlc.Timestamp{}
			}
			if err := m.TsMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TsMax", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TsMax == nil {
				m.TsMax = &hlc.Timestamp{}
			}
			if err := m.TsMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRocksdb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRocksdb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SSTUserPropertiesCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRocksdb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SSTUserPropertiesCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SSTUserPropertiesCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sst", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sst = append(m.Sst, SSTUserProperties{})
			if err := m.Sst[len(m.Sst)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRocksdb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRocksdb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HistogramData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRocksdb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HistogramData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HistogramData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mean", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Mean = float64(math.Float64frombits(v))
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field P50", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.P50 = float64(math.Float64frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field P95", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.P95 = float64(math.Float64frombits(v))
		case 4:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field P99", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.P99 = float64(math.Float64frombits(v))
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Max", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Max = float64(math.Float64frombits(v))
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sum", wireType)
			}
			m.Sum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sum |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRocksdb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRocksdb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TickersAndHistograms) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRocksdb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TickersAndHistograms: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickersAndHistograms: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tickers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tickers == nil {
				m.Tickers = make(map[string]uint64)
			}
			var mapkey string
			var mapvalue uint64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRocksdb
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRocksdb
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthRocksdb
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRocksdb
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipRocksdb(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthRocksdb
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tickers[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Histograms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRocksdb
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Histograms == nil {
				m.Histograms = make(map[string]HistogramData)
			}
			var mapkey string
			mapvalue := &HistogramData{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowRocksdb
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRocksdb
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthRocksdb
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowRocksdb
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthRocksdb
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthRocksdb
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &HistogramData{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipRocksdb(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthRocksdb
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Histograms[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRocksdb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRocksdb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRocksdb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRocksdb
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRocksdb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRocksdb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRocksdb
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRocksdb(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRocksdb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRocksdb   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("storage/enginepb/rocksdb.proto", fileDescriptor_rocksdb_e0dede928993ad0c)
}

var fileDescriptor_rocksdb_e0dede928993ad0c = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xde, 0x49, 0xb6, 0xa9, 0x9d, 0x28, 0xea, 0x92, 0xc3, 0xb2, 0xe2, 0x18, 0x72, 0x8a, 0x82,
	0xbb, 0x12, 0x0d, 0x98, 0x82, 0xa8, 0xd5, 0x82, 0x17, 0x41, 0xa6, 0x11, 0xc1, 0x8b, 0x4c, 0xa6,
	0xe3, 0xee, 0x90, 0xdd, 0x99, 0x65, 0x66, 0x22, 0xa9, 0xbf, 0x42, 0xf0, 0x4f, 0xe5, 0xd8, 0x63,
	0x4f, 0xa2, 0xc9, 0xef, 0x10, 0x64, 0x66, 0xb3, 0x6d, 0xb4, 0xb5, 0xe8, 0xed, 0x7b, 0x1f, 0xef,
	0xfb, 0xde, 0x7b, 0xdf, 0xee, 0x40, 0xa4, 0x8d, 0x54, 0x24, 0x65, 0x09, 0x13, 0x29, 0x17, 0xac,
	0x9c, 0x24, 0x4a, 0xd2, 0xa9, 0x3e, 0x9c, 0xc4, 0xa5, 0x92, 0x46, 0x06, 0x11, 0x95, 0x74, 0xaa,
	0x24, 0xa1, 0x59, 0xbc, 0xee, 0x8c, 0xeb, 0xce, 0x28, 0x9c, 0x19, 0x9e, 0x27, 0x59, 0x4e, 0x13,
	0xc3, 0x0b, 0xa6, 0x0d, 0x29, 0xca, 0x4a, 0x15, 0x75, 0x52, 0x99, 0x4a, 0x07, 0x13, 0x8b, 0x2a,
	0xb6, 0xf7, 0x15, 0xc0, 0x9b, 0x07, 0x07, 0xe3, 0xb7, 0x9a, 0xa9, 0x37, 0x4a, 0x96, 0x4c, 0x19,
	0xce, 0x74, 0x10, 0x40, 0xbf, 0x24, 0x26, 0x0b, 0x41, 0x17, 0xf4, 0x77, 0xb0, 0xc3, 0xc1, 0x23,
	0xd8, 0x32, 0xfa, 0x43, 0xc1, 0x45, 0xd8, 0xe8, 0x82, 0x7e, 0x7b, 0x70, 0x3b, 0x3e, 0x5b, 0xc3,
	0x0e, 0x8d, 0xb3, 0x9c, 0xc6, 0xe3, 0x7a, 0x28, 0xde, 0x32, 0xfa, 0x35, 0x17, 0xb5, 0x8a, 0xcc,
	0xc3, 0xe6, 0xbf, 0xaa, 0xc8, 0xbc, 0xf7, 0x19, 0xde, 0x3a, 0xb7, 0xd4, 0x0b, 0x99, 0xe7, 0x8c,
	0x1a, 0x2e, 0x45, 0xb0, 0x0f, 0x9b, 0x5a, 0x9b, 0x10, 0x74, 0x9b, 0xfd, 0xf6, 0xe0, 0x7e, 0xfc,
	0xf7, 0x38, 0xe2, 0x73, 0x2e, 0x7b, 0xfe, 0xe2, 0xdb, 0x1d, 0x0f, 0x5b, 0x7d, 0xd0, 0x81, 0x5b,
	0x4c, 0x29, 0xa9, 0xdc, 0x41, 0x3b, 0xb8, 0x2a, 0x6c, 0x22, 0xd7, 0x5e, 0x71, 0x6d, 0x64, 0xaa,
	0x48, 0xf1, 0x92, 0x18, 0x62, 0xd3, 0x28, 0x18, 0x11, 0x2e, 0x0d, 0x80, 0x1d, 0x0e, 0x6e, 0xc0,
	0x66, 0x39, 0x7c, 0xe0, 0x94, 0x00, 0x5b, 0xe8, 0x98, 0xd1, 0xd0, 0x9d, 0x69, 0x99, 0xd1, 0xb0,
	0x62, 0x46, 0xa1, 0x5f, 0x33, 0x23, 0xcb, 0xd8, 0x28, 0xb6, 0x2a, 0xa6, 0x20, 0x73, 0xbb, 0x03,
	0x95, 0x33, 0x61, 0xc2, 0x56, 0x17, 0xf4, 0x7d, 0x5c, 0x15, 0xb6, 0x4f, 0xcf, 0x8a, 0x70, 0xdb,
	0x71, 0x16, 0xf6, 0x7e, 0x36, 0x60, 0x67, 0xcc, 0xe9, 0x94, 0x29, 0xfd, 0x5c, 0x1c, 0x9e, 0xee,
	0xa7, 0x83, 0x77, 0x70, 0xdb, 0x54, 0xfc, 0x3a, 0x8f, 0x27, 0x97, 0xe5, 0x71, 0x91, 0x45, 0x4d,
	0xee, 0x0b, 0xa3, 0x8e, 0x70, 0xed, 0x16, 0x7c, 0x84, 0x30, 0x3b, 0xed, 0x09, 0x1b, 0xce, 0xfb,
	0xd9, 0x7f, 0x7b, 0x9f, 0x41, 0x67, 0xbf, 0x8e, 0x7f, 0xc3, 0x39, 0xda, 0x85, 0x57, 0x37, 0x17,
	0xb0, 0xb7, 0x4f, 0xd9, 0xd1, 0xfa, 0xd7, 0xb3, 0xd0, 0x66, 0xf4, 0x89, 0xe4, 0x33, 0xe6, 0xd2,
	0xf6, 0x71, 0x55, 0xec, 0x36, 0x1e, 0x83, 0x28, 0x83, 0xd7, 0xff, 0x18, 0x70, 0x81, 0xfc, 0xe9,
	0xa6, 0xbc, 0x3d, 0xb8, 0x7b, 0xd9, 0x0d, 0xbf, 0x7d, 0xf8, 0x8d, 0x49, 0x7b, 0xf7, 0x16, 0x3f,
	0x90, 0xb7, 0x58, 0x22, 0x70, 0xbc, 0x44, 0xe0, 0x64, 0x89, 0xc0, 0xf7, 0x25, 0x02, 0x5f, 0x56,
	0xc8, 0x3b, 0x5e, 0x21, 0xef, 0x64, 0x85, 0xbc, 0xf7, 0x57, 0x6a, 0x93, 0x49, 0xcb, 0x3d, 0xad,
	0x87, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0x8e, 0x19, 0xa2, 0xc8, 0x03, 0x00, 0x00,
}
