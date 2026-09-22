package neoqueries

import (
	"strings"
)

type QueryBuilder struct {
	*Builder
	parts []Buildable
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		Builder: NewBuilder(),
	}
}

func (qb *QueryBuilder) addPart(part Buildable) {
	qb.parts = append(qb.parts, part)
}

func (qb *QueryBuilder) Create(elements ...Element) *QueryBuilder {
	qb.addPart(newCreateBuilder(qb.Builder, elements))
	return qb
}

func (qb *QueryBuilder) Match(element ...Element) *QueryBuilder {
	qb.addPart(newMatchBuilder(qb.Builder, element))
	return qb
}

func (qb *QueryBuilder) Return(refs ...QueryRef) *QueryBuilder {
	qb.addPart(newReturnBuilder(qb.Builder, refs...))
	return qb
}

func (qb *QueryBuilder) Where(condition ConditionBuilderInterface) *QueryBuilder {
	qb.addPart(newWhereBuilder(qb.Builder, condition))
	return qb
}

func (qb *QueryBuilder) Build() (string, Params) {
	var b strings.Builder

	for _, part := range qb.parts {
		query := part.build()
		b.WriteString(query)
		b.WriteByte('\n')
	}

	return b.String(), qb.GetParams()
}
