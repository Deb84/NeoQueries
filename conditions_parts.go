package neoqueries

type conditionPart interface {
	build(*Builder) string
}

type refConditionPart struct {
	ref *UnbuiltRef
}

func (p *refConditionPart) build(builder *Builder) string {
	return p.ref.buildRef(builder).String()
}

type propConditionPart struct {
	prop *UnbuiltProp
}

func (p *propConditionPart) build(builder *Builder) string {
	return p.prop.buildRef(builder).String()
}

type nestedConditionPart struct {
	condition *ConditionBuilder
}

func (p *nestedConditionPart) build(builder *Builder) string {
	return "(" + p.condition.buildCondition(builder) + ")"
}

type rawPart string

func (p rawPart) build(*Builder) string {
	return string(p)
}

type anyPart struct {
	value any
}

func (p *anyPart) build(builder *Builder) string {
	return "$" + builder.nextValueRef(p.value).String()
}

type listPart[V any] struct {
	list *List[V]
}

func (p listPart[V]) build(builder *Builder) string {
	return "$" + builder.nextListRef(p.list).String()
}
