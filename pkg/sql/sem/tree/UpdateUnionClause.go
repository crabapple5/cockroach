package tree

// UpdateUnionClause represents a INSERT UNION SELECT statement.
type UpdateUnionClause struct {
	Type  UnionType
	Left  *Update
	Right *Update
	All   bool
}

// Format implements the NodeFormatter interface.
func (node *UpdateUnionClause) Format(ctx *FmtCtx) {
	//ctx.FormatNode(node.Left)
	ctx.WriteByte(' ')
	ctx.WriteString(node.Type.String())
	if node.All {
		ctx.WriteString(" ALL")
	}
	ctx.WriteByte(' ')
	ctx.FormatNode(node.Right)
}
