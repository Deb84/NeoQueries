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

type patternRefPart[R Reference] struct {
	ref *UnbuiltRef[R]
}

func (p *patternRefPart[R]) build(context *BuildContext) PatternString {
	return PatternString("(" + p.ref.buildRef(context).String() + ")")
}
