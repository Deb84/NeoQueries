package neoqueries

import "strings"

type PatternString string

func (p PatternString) string() string {
	return string(p)
}

type PatternInterface interface {
	build(context *BuildContext) PatternString
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
