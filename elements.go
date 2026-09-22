package queries

import "NeoQueries"

type Element interface {
	setBuilder(*NeoQueries.Builder)
	build() string
	GetParams() NeoQueries.Params
	GetRef() Ref
	element()
}
type ElementBuilder struct {
	*NeoQueries.Builder
	refs  NeoQueries.Refs
	props *NeoQueries.Props
}

func newElementBuilder() *ElementBuilder {
	props := make(NeoQueries.Props)

	return &ElementBuilder{
		refs:  NeoQueries.newRefs(),
		props: &props,
	}
}

func (eb *ElementBuilder) GetRef() Ref {
	return eb.refs.element
}
