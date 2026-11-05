package neoqueries

import (
	"fmt"
	"reflect"
	"strings"
)

type MatchBuilder struct {
	patterns []PatternBuildInterface
}

func newMatchBuilder[T ValidPattern](pattern T, patterns []PatternBuildInterface) *MatchBuilder {
	// Return an error if a pattern from patterns is not a ValidPattern
	for _, p := range patterns {
		if !IsAValidPattern(p) {
			panic(fmt.Sprintf("%s is not a valid Pattern", reflect.TypeOf(p)))
		}
	}

	patterns = append([]PatternBuildInterface{pattern}, patterns...)

	return &MatchBuilder{
		patterns: patterns,
	}
}

func NewMatchBuilder[T ValidPattern](pattern T, patterns ...PatternBuildInterface) *MatchBuilder {
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
