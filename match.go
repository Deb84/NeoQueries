package neoqueries

import (
	"strings"
)

type MatchBuilder struct {
	patterns []PatternInterface
}

func newMatchBuilder(pattern PatternInterface, patterns []PatternInterface) *MatchBuilder {
	patterns = append([]PatternInterface{pattern}, patterns...)

	return &MatchBuilder{
		patterns: patterns,
	}
}

func NewMatchBuilder(pattern PatternInterface, patterns ...PatternInterface) *MatchBuilder {
	return newMatchBuilder(pattern, patterns)
}

func (mb *MatchBuilder) build(builder *Builder) string {
	query := `MATCH`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, pattern := range mb.patterns {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(pattern.build(builder).string())
	}

	return b.String()
}
