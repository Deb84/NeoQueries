package neoqueries

import "strings"

type PatternString string

func (p PatternString) string() string {
	return string(p)
}

type PatternInterface interface {
	build(builder *Builder) PatternString
}

type patternPartInterface interface {
	build(builder *Builder) PatternString
}

type patternPart struct {
	props    *Props
	propsRef Ref
	tokenRef map[string]Ref
}

func newPatternPart() *patternPart {
	props := make(Props)
	return &patternPart{
		props:    &props,
		tokenRef: make(map[string]Ref),
	}
}

type Pattern struct {
	parts   []patternPartInterface
	builder *Builder
}

func NewPattern() *StartPatternState {
	p := &Pattern{}
	return &StartPatternState{
		pattern: p,
	}
}

func (p *Pattern) addPart(part patternPartInterface) {
	p.parts = append(p.parts, part)
}

func (p *Pattern) build(builder *Builder) PatternString {
	var b strings.Builder

	for _, part := range p.parts {
		b.WriteString(part.build(builder).string())
	}

	return PatternString(b.String())
}

func (p *Pattern) setBuilder(builder *Builder) {
	p.builder = builder
}

type StartPatternState struct {
	pattern *Pattern
}

func (p *StartPatternState) Node(node *NodePattern) *CompletePatternState {
	p.pattern.addPart(node)
	return newCompletePatternState(p.pattern)
}

type NodePatternState struct {
	pattern *Pattern
}

func (p *NodePatternState) Relation(relation *RelationPattern) *RelationPatternState {
	p.pattern.addPart(relation)
	return &RelationPatternState{pattern: p.pattern}
}

type RelationPatternState struct {
	pattern *Pattern
}

func (r *RelationPatternState) Node(node *NodePattern) *CompletePatternState {
	r.pattern.addPart(node)
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
