package neoqueries

import (
	"strings"
)

type ConditionBuilderInterface interface {
	buildCondition(builder *Builder) string
}

type ConditionBuilder struct {
	builder *Builder
	parts   []conditionPart
}

func NewConditionBuilder() *ConditionStart {
	cb := &ConditionBuilder{}
	return &ConditionStart{builder: cb}
}

func (cb *ConditionBuilder) addPart(part conditionPart) {
	cb.parts = append(cb.parts, part)
}

func (cb *ConditionBuilder) buildCondition(builder *Builder) string {
	cb.builder = builder

	var b strings.Builder

	for i, part := range cb.parts {
		if i > 0 {
			b.WriteByte(' ')
		}

		b.WriteString(part.build(cb))
	}

	return b.String()
}
