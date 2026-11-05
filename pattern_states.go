package neoqueries

type StartPatternState struct {
	pattern *Pattern
}

func (p *StartPatternState) Node(node *NodePattern) *CompletePatternState {
	p.pattern.addPart(node)
	return newCompletePatternState(p.pattern)
}

func (p *StartPatternState) NodeRef(ref *UnbuiltRef[NodeRef]) *CompletePatternState {
	p.pattern.addPart(&PatternRefPart[NodeRef]{ref})
	return newCompletePatternState(p.pattern)
}

type NodePatternState struct {
	pattern *Pattern
}

func (p *NodePatternState) Relation(relation *RelationPattern) *RelationPatternState {
	p.pattern.addPart(&RelationPatternPart[*RelationPattern]{
		pattern:   relation,
		direction: relation.direction,
	})
	return &RelationPatternState{pattern: p.pattern}
}

func (p *NodePatternState) RelationRef(ref *UnbuiltRef[RelationRef], direction RelationDirectionString) *RelationPatternState {
	part := &RelationPatternPart[*RelationRefPattern]{
		pattern:   &RelationRefPattern{ref},
		direction: direction,
	}

	p.pattern.addPart(part)
	return &RelationPatternState{pattern: p.pattern}
}

type RelationPatternState struct {
	pattern *Pattern
}

func (p *RelationPatternState) Node(node *NodePattern) *CompletePatternState {
	p.pattern.addPart(node)
	return newCompletePatternState(p.pattern)
}

func (p *RelationPatternState) NodeRef(ref *UnbuiltRef[NodeRef]) *CompletePatternState {
	p.pattern.addPart(&PatternRefPart[NodeRef]{ref})
	return newCompletePatternState(p.pattern)
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
