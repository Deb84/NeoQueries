package neoqueries

type Node struct {
	*Element
	ref Ref
}

func NewNode() *Node {
	return &Node{
		Element: newElement(),
	}
}

func (n *Node) getRef() Ref {
	return n.ref
}

func (n *Node) GetRef() *UnbuiltRef {
	return newUnbuiltRef(n)
}

func (n *Node) buildElement() Ref {
	n.ref = n.builder.nextNodeRef(n)
	return n.ref
}
