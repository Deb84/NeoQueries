package neoqueries

import (
	"strings"
)

type MatchBuilder struct {
	*Builder
	elements []Element
}

func newMatchBuilder(builder *Builder, element Element, elements []Element) *MatchBuilder {
	elements = append(elements, element)

	for _, e := range elements {
		e.setBuilder(builder)
	}

	return &MatchBuilder{
		Builder:  builder,
		elements: elements,
	}
}

func NewMatchBuilder(element Element, elements ...Element) *MatchBuilder {
	return newMatchBuilder(NewBuilder(), element, elements)
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
