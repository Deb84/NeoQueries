package neoqueries

type ElementInterface interface {
	setBuilder(*Builder)
	buildElement() Ref
	getRef() Ref
	GetRef() *UnbuiltRef
}
type Element struct {
	builder *Builder
}

func newElement() *Element {
	return &Element{}
}

func (e *Element) setBuilder(builder *Builder) {
	e.builder = builder
}
