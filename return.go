package neoqueries

import (
	"strings"
)

type ReturnBuilderBuild[T any] func(*BuildContext, T) string

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

func newReturnElementBuilder[R Reference, T Element[R]](ref T, refs []T) *ReturnBuilder[T] {
	return newBaseReturnBuilder(returnElementBuild[R], ref, refs)
}

func NewReturnBuilder[U PropOrRef, T UnbuiltPropOrVar[U]](ref T, refs ...T) *ReturnBuilder[T] {
	return newReturnBuilder(ref, refs)
}

func NewReturnElementBuilder[R Reference, T Element[R]](element T, elements ...T) *ReturnBuilder[T] {
	return newReturnElementBuilder(element, elements)
}

func (b *ReturnBuilder[T]) build(context *BuildContext) string {
	query := `RETURN`
	var s strings.Builder

	s.WriteString(query)
	s.WriteByte(' ')

	for i, obj := range b.objs {
		if i > 0 {
			s.WriteByte(',')
		}
		s.WriteString(b.buildFunc(context, obj))
	}

	return s.String()
}

func returnBuild[U PropOrRef, T UnbuiltPropOrVar[U]](context *BuildContext, ref T) string {
	return ref.buildRef(context).String()
}

func returnElementBuild[R Reference, T Element[R]](b *BuildContext, element T) string {
	return element.buildElement(b).String()
}
