package NeoQueries

import (
	"strings"
)

type MatchBuilder struct {
	*Builder
	elements []Element
}

func newMatchBuilder(builder *Builder, elements []Element) *MatchBuilder {
	for _, element := range elements {
		element.setBuilder(builder)
	}

	return &MatchBuilder{
		Builder:  builder,
		elements: elements,
	}
}

func NewMatchBuilder(elements ...Element) *MatchBuilder {
	return newMatchBuilder(NewBuilder(), elements)
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

func (mb *MatchBuilder) setBuilder(builder *Builder) {
	mb.Builder = builder
}
