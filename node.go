package neoqueries

type Node struct {
	ref Ref
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) getRef() Ref {
	return n.ref
}

func (n *Node) GetRef() *UnbuiltRef {
	return newUnbuiltRef(n)
}

func (n *Node) buildElement(builder *Builder) Ref {
	n.ref = builder.nextNodeRef(n)
	return n.ref
}
