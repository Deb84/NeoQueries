package neoqueries

import "strings"

type PatternString string

func (p PatternString) string() string {
	return string(p)
}

type PatternBuildInterface interface {
	build(context *BuildContext) PatternString
}

// ValidPattern defines the valid Pattern types for the clauses
type ValidPattern interface {
	*NodePattern | *PatternRefPart[NodeRef] | *PatternRefPart[RelationRef] | *CompletePatternState
	build(context *BuildContext) PatternString
}

func IsAValidPattern(pattern PatternBuildInterface) bool {
	switch pattern.(type) {
	case *NodePattern, *PatternRefPart[NodeRef], *PatternRefPart[RelationRef], *CompletePatternState:
		return true
	default:
		return false
	}
}

// Pattern implements a neo4j pattern
type Pattern struct {
	parts []patternPartInterface
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

func (p *Pattern) build(context *BuildContext) PatternString {
	var b strings.Builder

	for _, part := range p.parts {
		b.WriteString(part.build(context).string())
	}

	return PatternString(b.String())
}
