package neoqueries

import (
	"strings"
)

type WhereBuilder struct {
	condition ConditionBuilderInterface
}

func newWhereBuilder(condition ConditionBuilderInterface) *WhereBuilder {
	return &WhereBuilder{
		condition: condition,
	}
}

func NewWhereBuilder(condition ConditionBuilderInterface) *WhereBuilder {
	return newWhereBuilder(condition)
}

func (wb *WhereBuilder) build(builder *Builder) string {
	query := `WHERE`

	var b strings.Builder

	b.WriteString(query)
	b.WriteByte(' ')
	b.WriteString(wb.condition.buildCondition(builder))

	return b.String()
}
