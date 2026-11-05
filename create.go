package neoqueries

import (
	"strings"
)

type CreateBuilder struct {
	patterns []PatternBuildInterface
}

func newCreateBuilder[S ValidPatternString, T ValidPattern[S]](pattern T, patterns []PatternBuildInterface) *CreateBuilder {
	patterns = append([]PatternBuildInterface{PatternAdapter[S, T]{pattern}}, patterns...)

	return &CreateBuilder{
		patterns: patterns,
	}
}

func NewCreateBuilder[S ValidPatternString, T ValidPattern[S]](element T, elements ...PatternBuildInterface) *CreateBuilder {
	return newCreateBuilder(element, elements)
}

func (cb *CreateBuilder) build(context *BuildContext) string {
	query := `CREATE`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, pattern := range cb.patterns {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(pattern.build(context).string())
	}

	return b.String()
}
