package neoqueries

type conditionPart interface {
	build(*ConditionBuilder) string
}

type refConditionPart struct {
	ref *UnbuiltRef
}

func (p *refConditionPart) build(cb *ConditionBuilder) string {
	return p.ref.buildRef(cb.builder).String()
}

type propConditionPart struct {
	prop *UnbuiltProp
}

func (p *propConditionPart) build(cb *ConditionBuilder) string {
	return p.prop.buildRef(cb.builder).String()
}

type nestedConditionPart struct {
	condition *ConditionBuilder
}

func (p *nestedConditionPart) build(cb *ConditionBuilder) string {
	return "(" + p.condition.buildCondition(cb.builder) + ")"
}

type rawPart string

func (p rawPart) build(*ConditionBuilder) string {
	return string(p)
}

type anyPart struct {
	value any
}

func (p *anyPart) build(cb *ConditionBuilder) string {
	return "$" + cb.builder.nextValueRef(p.value).String()
}

type listPart[V any] struct {
	list *List[V]
}

func (p listPart[V]) build(cb *ConditionBuilder) string {
	return "$" + cb.builder.nextListRef(p.list).String()
}
