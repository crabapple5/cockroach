package tree

// LazySelectUnionClause represents a INSERT UNION SELECT statement.
type LazySelectUnionClause struct {
	Type  UnionType
	Left  *Insert
	Right *Select
	All   bool
}

// Format implements the NodeFormatter interface.
func (node *LazySelectUnionClause) Format(ctx *FmtCtx) {
	//ctx.FormatNode(node.Left)
	ctx.WriteByte(' ')
	ctx.WriteString(node.Type.String())
	if node.All {
		ctx.WriteString(" ALL")
	}
	ctx.WriteByte(' ')
	ctx.FormatNode(node.Right)
}
