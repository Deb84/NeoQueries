package NeoQueries

import (
	"strings"
)

type ConditionBuilderInterface interface {
	buildCondition() string
	setBuilder(builder *Builder)
}

type ConditionBuilder struct {
	*Builder
	parts []conditionPart
}

func NewConditionBuilder() *ConditionStart {
	cb := &ConditionBuilder{}
	return &ConditionStart{builder: cb}
}

func (cb *ConditionBuilder) addPart(part conditionPart) {
	cb.parts = append(cb.parts, part)
}

func (cb *ConditionBuilder) buildCondition() string {
	var b strings.Builder

	for i, part := range cb.parts {
		if i > 0 {
			b.WriteByte(' ')
		}

		b.WriteString(part.build(cb))
	}

	return b.String()
}

func (cb *ConditionBuilder) setBuilder(builder *Builder) {
	cb.Builder = builder
}
