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

func (qb *QueryBuilder) Create(pattern PatternInterface, patterns ...PatternInterface) *QueryBuilder {
	qb.addPart(newCreateBuilder(qb.Builder, pattern, patterns))
	return qb
}

func (qb *QueryBuilder) Delete(element ElementInterface, elements ...ElementInterface) *QueryBuilder {
	qb.addPart(newDeleteBuilder(qb.Builder, element, elements))
	return qb
}

func (qb *QueryBuilder) DetachDelete(element ElementInterface, elements ...ElementInterface) *QueryBuilder {
	qb.addPart(newDetachDeleteBuilder(qb.Builder, element, elements))
	return qb
}

func (qb *QueryBuilder) Match(pattern PatternInterface, patterns ...PatternInterface) *QueryBuilder {
	qb.addPart(newMatchBuilder(qb.Builder, pattern, patterns))
	return qb
}

func (qb *QueryBuilder) Return[U PropOrRef, T UnbuiltPropOrVar[U]](ref T, refs ...T) *QueryBuilder {
	qb.addPart(newReturnBuilder(qb.Builder, ref, refs))
	return qb
}

func (qb *QueryBuilder) ReturnElement(element ElementInterface, elements ...ElementInterface) *QueryBuilder {
	qb.addPart(newReturnElementBuilder(qb.Builder, element, elements))
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
