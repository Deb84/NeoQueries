package neoqueries

import (
	"errors"
)

type Element interface {
	setBuilder(*Builder)
	build() string
	GetParams() Params
	GetRef() (Ref, error)
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

func (eb *ElementBuilder) GetRef() (Ref, error) {
	if eb.refs.element != "" {
		return eb.refs.element, nil
	}

	return "", errors.New("this element doesn't have reference, element need to be built")
}
