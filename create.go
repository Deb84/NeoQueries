package neoqueries

import (
	"strings"
)

type CreateBuilder struct {
	*Builder
	patterns []PatternInterface
}

func newCreateBuilder(builder *Builder, pattern PatternInterface, patterns []PatternInterface) *CreateBuilder {
	patterns = append([]PatternInterface{pattern}, patterns...)

	for _, e := range patterns {
		e.setBuilder(builder)
	}

	return &CreateBuilder{
		Builder:  builder,
		patterns: patterns,
	}
}

func NewCreateBuilder(element PatternInterface, elements ...PatternInterface) *CreateBuilder {
	return newCreateBuilder(NewBuilder(), element, elements)
}

func (cb *CreateBuilder) build() string {
	query := `CREATE`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, pattern := range cb.patterns {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(pattern.build().string())
	}

	return b.String()
}
