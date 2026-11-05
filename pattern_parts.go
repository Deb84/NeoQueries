package neoqueries

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

// PatternRefPart allows an UnbuiltRef to be used to refer to an Element
type PatternRefPart[R Reference] struct {
	ref *UnbuiltRef[R]
}

func (p *PatternRefPart[R]) build(context *BuildContext) PatternString {
	ref := "(" + p.ref.buildRef(context).String() + ")"
	return PatternString(ref)
}
