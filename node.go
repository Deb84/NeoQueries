package neoqueries

type Node struct {
	ref NodeRef
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) GetRef() *UnbuiltRef[NodeRef] {
	return newUnbuiltRef(n)
}

func (n *Node) buildElement(context *BuildContext) NodeRef {
	return context.ensureNodeRef(n)
}
