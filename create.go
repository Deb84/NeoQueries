package queries

import (
	"NeoQueries"
	"strings"
)

type CreateBuilder struct {
	*NeoQueries.Builder
	elements []Element
}

func newCreateBuilder(builder *NeoQueries.Builder, elements []Element) *CreateBuilder {
	for _, element := range elements {
		element.setBuilder(builder)
	}

	return &CreateBuilder{
		NeoQueries.Builder: builder,
		elements:           elements,
	}
}

func NewCreateBuilder(elements ...Element) *CreateBuilder {
	return newCreateBuilder(NeoQueries.NewBuilder(), elements)
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
