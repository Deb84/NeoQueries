package neoqueries

type StartPatternState struct {
	pattern *Pattern
}

func (p *StartPatternState) Node(node *NodePattern) *CompletePatternState {
	p.pattern.addPart(patternPartAdapter[PatternNodeString]{part: node})
	return newCompletePatternState(p.pattern)
}

func (p *StartPatternState) NodeRef(ref *UnbuiltRef[NodeRef]) *CompletePatternState {
	p.pattern.addPart(patternPartAdapter[PatternRefString]{part: &PatternRefPart[NodeRef]{ref}})
	return newCompletePatternState(p.pattern)
}

type NodePatternState struct {
	pattern *Pattern
}

func (p *NodePatternState) Relation(relation *RelationPattern) *RelationPatternState {
	p.pattern.addPart(patternPartAdapter[PatternRelationString]{part: relation})
	return &RelationPatternState{pattern: p.pattern}
}

func (p *NodePatternState) RelationRef(ref *UnbuiltRef[RelationRef]) *RelationPatternState {
	p.pattern.addPart(patternPartAdapter[PatternRefString]{part: &PatternRefPart[RelationRef]{ref}})
	return &RelationPatternState{pattern: p.pattern}
}

type RelationPatternState struct {
	pattern *Pattern
}

func (p *RelationPatternState) Node(node *NodePattern) *CompletePatternState {
	p.pattern.addPart(patternPartAdapter[PatternNodeString]{part: node})
	return newCompletePatternState(p.pattern)
}

func (p *RelationPatternState) NodeRef(ref *UnbuiltRef[NodeRef]) *CompletePatternState {
	p.pattern.addPart(patternPartAdapter[PatternRefString]{part: &PatternRefPart[NodeRef]{ref}})
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
