package neoqueries

import "strings"

type PatternString string

func (p PatternString) string() string {
	return string(p)
}

type PatternNodeString struct {
	PatternString
}
type PatternRelationString struct {
	PatternString
}
type PatternRefString[T Reference] struct {
	PatternString
}

type patternStinger interface {
	string() string
}

type PatternBuildInterface interface {
	build(context *BuildContext) patternStinger
}

// ValidPattern defines the valid Pattern types for the clauses
type ValidPattern[T patternStinger] interface {
	*NodePattern | *PatternRefPart[NodeRef] | *PatternRefPart[RelationRef] | *CompletePatternState
	build(context *BuildContext) T
}

// PatternAdapter convert a patternStinger from a ValidPattern to a patternStinger
type PatternAdapter[S patternStinger, T ValidPattern[S]] struct {
	pattern T
}

func (p PatternAdapter[S, T]) build(context *BuildContext) patternStinger {
	return p.pattern.build(context)
}

// Pattern implements a neo4j pattern
type Pattern struct {
	parts []patternPartBuildInterface
}

func NewPattern() *StartPatternState {
	p := &Pattern{}
	return &StartPatternState{
		pattern: p,
	}
}

func (p *Pattern) addPart(part patternPartBuildInterface) {
	p.parts = append(p.parts, part)
}

func (p *Pattern) build(context *BuildContext) PatternString {
	var b strings.Builder

	for _, part := range p.parts {
		b.WriteString(part.build(context).string())
	}

	return PatternString(b.String())
}
