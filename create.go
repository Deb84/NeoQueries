package neoqueries

import (
	"strings"
)

type CreateBuilder struct {
	patterns []PatternInterface
}

func newCreateBuilder(pattern PatternInterface, patterns []PatternInterface) *CreateBuilder {
	patterns = append([]PatternInterface{pattern}, patterns...)

	return &CreateBuilder{
		patterns: patterns,
	}
}

func NewCreateBuilder(element PatternInterface, elements ...PatternInterface) *CreateBuilder {
	return newCreateBuilder(element, elements)
}

func (cb *CreateBuilder) build(builder *Builder) string {
	query := `CREATE`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, pattern := range cb.patterns {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(pattern.build(builder).string())
	}

	return b.String()
}
