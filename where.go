package neoqueries

import (
	"strings"
)

type WhereBuilder struct {
	*Builder
	condition ConditionBuilderInterface
}

func newWhereBuilder(builder *Builder, condition ConditionBuilderInterface) *WhereBuilder {
	condition.setBuilder(builder)

	return &WhereBuilder{
		Builder:   builder,
		condition: condition,
	}
}

func NewWhereBuilder(condition ConditionBuilderInterface) *WhereBuilder {
	return newWhereBuilder(nil, condition)
}

func (wb *WhereBuilder) build() string {
	query := `WHERE`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')
	b.WriteString(wb.condition.buildCondition())

	return b.String()
}

func (wb *WhereBuilder) setBuilder(builder *Builder) {
	wb.Builder = builder
}
