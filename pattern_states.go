package neoqueries

type StartPatternState struct {
	pattern *Pattern
}

func (p *StartPatternState) Node(node *NodePattern) *CompletePatternState {
	p.pattern.addPart(node)
	return newCompletePatternState(p.pattern)
}

func (p *StartPatternState) NodeRef(ref *UnbuiltRef[NodeRef]) *CompletePatternState {
	p.pattern.addPart(&patternRefPart[NodeRef]{ref})
	return newCompletePatternState(p.pattern)
}

type NodePatternState struct {
	pattern *Pattern
}

func (p *NodePatternState) Relation(relation *RelationPattern) *RelationPatternState {
	p.pattern.addPart(relation)
	return &RelationPatternState{pattern: p.pattern}
}

func (p *NodePatternState) RelationRef(ref *UnbuiltRef[RelationRef]) *RelationPatternState {
	p.pattern.addPart(&patternRefPart[RelationRef]{ref})
	return &RelationPatternState{pattern: p.pattern}
}

type RelationPatternState struct {
	pattern *Pattern
}

func (r *RelationPatternState) Node(node *NodePattern) *CompletePatternState {
	r.pattern.addPart(node)
	return newCompletePatternState(r.pattern)
}

func (r *RelationPatternState) NodeRef(ref *UnbuiltRef[NodeRef]) *CompletePatternState {
	r.pattern.addPart(&patternRefPart[NodeRef]{ref})
	return newCompletePatternState(r.pattern)
}

type CompletePatternState struct {
	*Pattern
	*NodePatternState
}

func newCompletePatternState(pattern *Pattern) *CompletePatternState {
	return &CompletePatternState{
		Pattern:          pattern,
		NodePatternState: &NodePatternState{pattern},
	}
}
