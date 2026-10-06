package neoqueries

import "strings"

type DeleteBuilder struct {
	elements []ElementInterface
	query    string
}

func newBaseDeleteBuilder(query string, element ElementInterface, elements []ElementInterface) *DeleteBuilder {
	elements = append([]ElementInterface{element}, elements...)

	return &DeleteBuilder{
		elements: elements,
		query:    query,
	}
}

func newDeleteBuilder(element ElementInterface, elements []ElementInterface) *DeleteBuilder {
	return newBaseDeleteBuilder("DELETE", element, elements)
}

func newDetachDeleteBuilder(element ElementInterface, elements []ElementInterface) *DeleteBuilder {
	return newBaseDeleteBuilder("DETACH DELETE", element, elements)
}

func NewDeleteBuilder(element ElementInterface, elements ...ElementInterface) *DeleteBuilder {
	return newBaseDeleteBuilder("DELETE", element, elements)
}

func NewDetachDeleteBuilder(element ElementInterface, elements ...ElementInterface) *DeleteBuilder {
	return newBaseDeleteBuilder("DETACH DELETE", element, elements)
}

func (b *DeleteBuilder) build(builder *Builder) string {
	var s strings.Builder

	s.WriteString(b.query)
	s.WriteByte(' ')

	for i, element := range b.elements {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(element.buildElement(builder).String())
	}

	return s.String()
}
