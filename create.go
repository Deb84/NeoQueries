package neoqueries

import (
	"strings"
)

type CreateBuilder struct {
	*Builder
	elements []Element
}

func newCreateBuilder(builder *Builder, element Element, elements []Element) *CreateBuilder {
	elements = append([]Element{element}, elements...)

	for _, e := range elements {
		e.setBuilder(builder)
	}

	return &CreateBuilder{
		Builder:  builder,
		elements: elements,
	}
}

func NewCreateBuilder(element Element, elements ...Element) *CreateBuilder {
	return newCreateBuilder(NewBuilder(), element, elements)
}

func (cb *CreateBuilder) build() string {
	query := `CREATE`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, element := range cb.elements {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(element.build())
	}

	return b.String()
}
