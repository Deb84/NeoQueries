package neoqueries

type Node struct {
	ref Ref
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) GetRef() *UnbuiltRef {
	return newUnbuiltRef(n)
}

func (n *Node) buildElement(builder *Builder) Ref {
	n.ref = builder.ensureNodeRef(n)
	return n.ref
}
