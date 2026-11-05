package neoqueries

type patternPartInterface[T patternStinger] interface {
	build(context *BuildContext) T
}

type patternPartBuildInterface interface {
	build(context *BuildContext) patternStinger
}

type patternPartAdapter[T patternStinger] struct {
	part patternPartInterface[T]
}

func (p patternPartAdapter[T]) build(context *BuildContext) patternStinger {
	return p.part.build(context)
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

func (p *PatternRefPart[R]) build(context *BuildContext) PatternRefString {
	ref := "(" + p.ref.buildRef(context).String() + ")"
	return PatternRefString{PatternString(ref)}
}
