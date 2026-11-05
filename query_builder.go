package neoqueries

import (
	"strings"
)

type QueryBuilder struct {
	*BuildContext
	parts []Buildable
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		BuildContext: NewBuildContext(defaultConfig),
	}
}

func newQueryBuilderWithConfig(config Config) *QueryBuilder {
	return &QueryBuilder{
		BuildContext: NewBuildContext(config),
	}
}

func (qb *QueryBuilder) addPart(part Buildable) {
	qb.parts = append(qb.parts, part)
}

func (qb *QueryBuilder) Create[T ValidPattern](pattern T, patterns ...PatternBuildInterface) *QueryBuilder {
	qb.addPart(newCreateBuilder(pattern, patterns))
	return qb
}

func (qb *QueryBuilder) Delete[T Reference](element Element[T], elements ...Element[T]) *QueryBuilder {
	qb.addPart(newDeleteBuilder(element, elements))
	return qb
}

func (qb *QueryBuilder) DetachDelete[T Reference](element Element[T], elements ...Element[T]) *QueryBuilder {
	qb.addPart(newDetachDeleteBuilder(element, elements))
	return qb
}

func (qb *QueryBuilder) Match[T ValidPattern](pattern T, patterns ...PatternBuildInterface) *QueryBuilder {
	qb.addPart(newMatchBuilder(pattern, patterns))
	return qb
}

func (qb *QueryBuilder) Return[U PropOrRef, T UnbuiltPropOrVar[U]](ref T, refs ...T) *QueryBuilder {
	qb.addPart(newReturnBuilder(ref, refs))
	return qb
}

func (qb *QueryBuilder) ReturnElement[T Reference](element Element[T], elements ...Element[T]) *QueryBuilder {
	qb.addPart(newReturnElementBuilder(element, elements))
	return qb
}

func (qb *QueryBuilder) Where(condition ConditionBuilderInterface) *QueryBuilder {
	qb.addPart(newWhereBuilder(condition))
	return qb
}

func (qb *QueryBuilder) Build() (string, Params) {
	var b strings.Builder

	for _, part := range qb.parts {
		query := part.build(qb.BuildContext)
		b.WriteString(query)
		b.WriteByte('\n')
	}

	return b.String(), qb.GetParams()
}
