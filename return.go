package neoqueries

import (
	"strings"
)

type ReturnBuilderBuild[T any] func(*ReturnBuilder[T], T) string

type ReturnBuilder[T any] struct {
	*Builder
	objs      []T
	buildFunc ReturnBuilderBuild[T]
}

func newBaseReturnBuilder[T any](builder *Builder, fun ReturnBuilderBuild[T], obj T, objs []T) *ReturnBuilder[T] {
	objs = append([]T{obj}, objs...)

	return &ReturnBuilder[T]{
		Builder:   builder,
		objs:      objs,
		buildFunc: fun,
	}
}

func newReturnBuilder[T QueryRef](builder *Builder, ref T, refs []T) *ReturnBuilder[T] {
	return newBaseReturnBuilder(builder, returnBuild[T], ref, refs)
}

func newReturnElementBuilder[T Element](builder *Builder, ref T, refs []T) *ReturnBuilder[T] {
	return newBaseReturnBuilder(builder, returnElementBuild[T], ref, refs)
}

func NewReturnBuilder[T QueryRef](ref T, refs ...T) *ReturnBuilder[T] {
	return newReturnBuilder(nil, ref, refs)
}

func NewReturnElementBuilder[T Element](element T, elements ...T) *ReturnBuilder[T] {
	return newReturnElementBuilder(nil, element, elements)
}

func (b *ReturnBuilder[T]) build() string {
	query := `RETURN`
	var s strings.Builder

	s.WriteString(query)
	s.WriteByte(' ')

	for i, obj := range b.objs {
		if i > 0 {
			s.WriteByte(',')
		}
		s.WriteString(b.buildFunc(b, obj))
	}

	return s.String()
}

func (b *ReturnBuilder[T]) setBuilder(builder *Builder) {
	b.Builder = builder
}

func returnBuild[T QueryRef](b *ReturnBuilder[T], ref T) string {
	return ref.toRef(b.Builder).String()
}

func returnElementBuild[T Element](b *ReturnBuilder[T], element T) string {
	ref, _ := element.GetRef() // TODO: error handling
	return ref.String()
}
