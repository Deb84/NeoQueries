package queries

import (
	"NeoQueries"
	"strings"
)

type WhereBuilder struct {
	*NeoQueries.Builder
	condition ConditionBuilderInterface
}

func newWhereBuilder(builder *NeoQueries.Builder, condition ConditionBuilderInterface) *WhereBuilder {
	condition.setBuilder(builder)

	return &WhereBuilder{
		NeoQueries.Builder: builder,
		condition:          condition,
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

func (wb *WhereBuilder) setBuilder(builder *NeoQueries.Builder) {
	wb.Builder = builder
}
