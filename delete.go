package neoqueries

import "strings"

type DeleteBuilder[T Reference] struct {
	elements []Element[T]
	query    string
}

func newBaseDeleteBuilder[T Reference](query string, element Element[T], elements []Element[T]) *DeleteBuilder[T] {
	elements = append([]Element[T]{element}, elements...)

	return &DeleteBuilder[T]{
		elements: elements,
		query:    query,
	}
}

func newDeleteBuilder[T Reference](element Element[T], elements []Element[T]) *DeleteBuilder[T] {
	return newBaseDeleteBuilder("DELETE", element, elements)
}

func newDetachDeleteBuilder[T Reference](element Element[T], elements []Element[T]) *DeleteBuilder[T] {
	return newBaseDeleteBuilder("DETACH DELETE", element, elements)
}

func NewDeleteBuilder[T Reference](element Element[T], elements ...Element[T]) *DeleteBuilder[T] {
	return newBaseDeleteBuilder("DELETE", element, elements)
}

func NewDetachDeleteBuilder[T Reference](element Element[T], elements ...Element[T]) *DeleteBuilder[T] {
	return newBaseDeleteBuilder("DETACH DELETE", element, elements)
}

func (b *DeleteBuilder[T]) build(context *BuildContext) string {
	var s strings.Builder

	s.WriteString(b.query)
	s.WriteByte(' ')

	for i, element := range b.elements {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(element.buildElement(context).String())
	}

	return s.String()
}
