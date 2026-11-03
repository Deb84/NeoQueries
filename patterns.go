package neoqueries

import "strings"

type PatternString string

func (p PatternString) string() string {
	return string(p)
}

type PatternInterface interface {
	build(context *BuildContext) PatternString
}

type patternPartInterface interface {
	build(context *BuildContext) PatternString
}

type patternPart struct {
	props    *Props
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

type patternRefPart[R Reference] struct {
	ref *UnbuiltRef[R]
}

func (p *patternRefPart[R]) build(context *BuildContext) PatternString {
	return PatternString("(" + p.ref.buildRef(context).String() + ")")
}
