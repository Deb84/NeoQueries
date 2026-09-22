package NeoQueries

type Element interface {
	setBuilder(*Builder)
	build() string
	GetParams() Params
	GetRef() Ref
	element()
}
type ElementBuilder struct {
	*Builder
	refs  Refs
	props *Props
}

func newElementBuilder() *ElementBuilder {
	props := make(Props)

	return &ElementBuilder{
		refs:  newRefs(),
		props: &props,
	}
}

func (eb *ElementBuilder) GetRef() Ref {
	return eb.refs.element
}
