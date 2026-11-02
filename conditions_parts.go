package neoqueries

type conditionPart interface {
	build(*BuildContext) string
}

type refConditionPart struct {
	ref *UnbuiltRef
}

func (p *refConditionPart) build(context *BuildContext) string {
	return p.ref.buildRef(context).String()
}

type propConditionPart struct {
	prop *UnbuiltProp
}

func (p *propConditionPart) build(context *BuildContext) string {
	return p.prop.buildRef(context).String()
}

type nestedConditionPart struct {
	condition *ConditionBuilder
}

func (p *nestedConditionPart) build(context *BuildContext) string {
	return "(" + p.condition.buildCondition(context) + ")"
}

type rawPart string

func (p rawPart) build(*BuildContext) string {
	return string(p)
}

type listPart[V any] struct {
	list *List[V]
}

func (p listPart[V]) build(context *BuildContext) string {
	return "$" + context.ensureListRef(p.list).String()
}

type anyPart struct {
	value any
}

func (p *anyPart) build(context *BuildContext) string {
	return "$" + context.ensureValueRef(p.value).String()
}

type ptrPart[T any] struct {
	value *T
}

func (p *ptrPart[T]) build(context *BuildContext) string {
	return "$" + context.ensureValuePtrRef(p.value).String()
}
