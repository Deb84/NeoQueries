package neoqueries

import (
	"fmt"
	"reflect"
	"strings"
)

type CreateBuilder struct {
	patterns []PatternBuildInterface
}

func newCreateBuilder[T ValidPattern](pattern T, patterns []PatternBuildInterface) *CreateBuilder {
	// Return an error if a pattern from patterns is not a ValidPattern
	for _, p := range patterns {
		if !IsAValidPattern(p) {
			panic(fmt.Sprintf("%s is not a valid Pattern", reflect.TypeOf(p)))
		}
	}

	patterns = append([]PatternBuildInterface{pattern}, patterns...)

	return &CreateBuilder{
		patterns: patterns,
	}
}

func NewCreateBuilder[T ValidPattern](pattern T, patterns ...PatternBuildInterface) *CreateBuilder {
	return newCreateBuilder(pattern, patterns)
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
