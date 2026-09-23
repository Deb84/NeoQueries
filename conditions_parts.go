package neoqueries

type conditionPart interface {
	build(*ConditionBuilder) string
}

type refConditionPart struct {
	ref QueryRef
}

func (p *refConditionPart) build(cb *ConditionBuilder) string {
	return p.ref.toRef(cb.Builder).String()
}

type nestedConditionPart struct {
	condition ConditionBuilderInterface
}

func (p *nestedConditionPart) build(cb *ConditionBuilder) string {
	p.condition.setBuilder(cb.Builder)
	return "(" + p.condition.buildCondition() + ")"
}

type rawPart string

func (p rawPart) build(*ConditionBuilder) string {
	return string(p)
}

type anyPart struct {
	value any
}

func (p *anyPart) build(cb *ConditionBuilder) string {
	return cb.nextValueRef(p.value).String()
}

type listPart[V any] struct {
	list *List[V]
}

func (p listPart[V]) build(cb *ConditionBuilder) string {
	ref := cb.nextListRef(p.list)
	return "$" + ref.String()
}
