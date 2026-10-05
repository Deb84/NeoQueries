package neoqueries

import (
	"strings"
)

type MatchBuilder struct {
	*Builder
	patterns []PatternInterface
}

func newMatchBuilder(builder *Builder, pattern PatternInterface, patterns []PatternInterface) *MatchBuilder {
	patterns = append([]PatternInterface{pattern}, patterns...)

	for _, e := range patterns {
		e.setBuilder(builder)
	}

	return &MatchBuilder{
		Builder:  builder,
		patterns: patterns,
	}
}

func NewMatchBuilder(pattern PatternInterface, patterns ...PatternInterface) *MatchBuilder {
	return newMatchBuilder(nil, pattern, patterns)
}

func (mb *MatchBuilder) build() string {
	query := `MATCH`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, pattern := range mb.patterns {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(pattern.build().string())
	}

	return b.String()
}

func (mb *MatchBuilder) setBuilder(builder *Builder) {
	mb.Builder = builder
}
