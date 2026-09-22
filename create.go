package neoqueries

import (
	"strings"
)

type CreateBuilder struct {
	*Builder
	elements []Element
}

func newCreateBuilder(builder *Builder, elements []Element) *CreateBuilder {
	for _, element := range elements {
		element.setBuilder(builder)
	}

	return &CreateBuilder{
		Builder:  builder,
		elements: elements,
	}
}

func NewCreateBuilder(elements ...Element) *CreateBuilder {
	return newCreateBuilder(NewBuilder(), elements)
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
