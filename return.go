package neoqueries

import (
	"strings"
)

type ReturnBuilderBuild[T any] func(*Builder, T) string

type ReturnBuilder[T any] struct {
	objs      []T
	buildFunc ReturnBuilderBuild[T]
}

func newBaseReturnBuilder[T any](fun ReturnBuilderBuild[T], obj T, objs []T) *ReturnBuilder[T] {
	objs = append([]T{obj}, objs...)

	return &ReturnBuilder[T]{
		objs:      objs,
		buildFunc: fun,
	}
}

func newReturnBuilder[U PropOrRef, T UnbuiltPropOrVar[U]](ref T, refs []T) *ReturnBuilder[T] {
	return newBaseReturnBuilder(returnBuild[U, T], ref, refs)
}

func newReturnElementBuilder[T ElementInterface](ref T, refs []T) *ReturnBuilder[T] {
	return newBaseReturnBuilder(returnElementBuild[T], ref, refs)
}

func NewReturnBuilder[U PropOrRef, T UnbuiltPropOrVar[U]](ref T, refs ...T) *ReturnBuilder[T] {
	return newReturnBuilder(ref, refs)
}

func NewReturnElementBuilder[T ElementInterface](element T, elements ...T) *ReturnBuilder[T] {
	return newReturnElementBuilder(element, elements)
}

func (b *ReturnBuilder[T]) build(builder *Builder) string {
	query := `RETURN`
	var s strings.Builder

	s.WriteString(query)
	s.WriteByte(' ')

	for i, obj := range b.objs {
		if i > 0 {
			s.WriteByte(',')
		}
		s.WriteString(b.buildFunc(builder, obj))
	}

	return s.String()
}

func returnBuild[U PropOrRef, T UnbuiltPropOrVar[U]](builder *Builder, ref T) string {
	return ref.buildRef(builder).String()
}

func returnElementBuild[T ElementInterface](b *Builder, element T) string {
	return element.buildElement(b).String()
}
