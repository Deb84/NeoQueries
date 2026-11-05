package neoqueries

import (
	"strings"
)

type MatchBuilder struct {
	patterns []PatternBuildInterface
}

func newMatchBuilder[S ValidPatternString, T ValidPattern[S]](pattern T, patterns []PatternBuildInterface) *MatchBuilder {
	patterns = append([]PatternBuildInterface{PatternAdapter[S, T]{pattern}}, patterns...)

	return &MatchBuilder{
		patterns: patterns,
	}
}

func NewMatchBuilder[S ValidPatternString, T ValidPattern[S]](
	pattern T,
	patterns ...PatternBuildInterface,
) *MatchBuilder {
	return newMatchBuilder(pattern, patterns)
}

func (mb *MatchBuilder) build(context *BuildContext) string {
	query := `MATCH`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')

	for i, pattern := range mb.patterns {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(pattern.build(context).string())
	}

	return b.String()
}
