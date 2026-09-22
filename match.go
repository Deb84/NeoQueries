package queries

import (
	"NeoQueries"
	"strings"
)

type MatchBuilder struct {
	*NeoQueries.Builder
	elements []Element
}

func newMatchBuilder(builder *NeoQueries.Builder, elements []Element) *MatchBuilder {
	for _, element := range elements {
		element.setBuilder(builder)
	}

	return &MatchBuilder{
		NeoQueries.Builder: builder,
		elements:           elements,
	}
}

func NewMatchBuilder(elements ...Element) *MatchBuilder {
	return newMatchBuilder(NeoQueries.NewBuilder(), elements)
}

func (mb *MatchBuilder) build() string {
	query := `MATCH`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, element := range mb.elements {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(element.build())
	}

	return b.String()
}

func (mb *MatchBuilder) setBuilder(builder *NeoQueries.Builder) {
	mb.Builder = builder
}
