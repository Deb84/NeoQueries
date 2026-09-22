package NeoQueries

import (
	"fmt"
	"strings"
)

type NodeBuilder struct {
	*ElementBuilder
	labels      []string
	accessProps []string
	empty       bool
}

func NewNodeBuilder(empty bool) *NodeBuilder {
	return &NodeBuilder{
		ElementBuilder: newElementBuilder(),
		empty:          empty,
	}
}

func (nb *NodeBuilder) Label(label string) *NodeBuilder {
	nb.labels = append(nb.labels, label)
	return nb
}

func (nb *NodeBuilder) Props(props Props) *NodeBuilder {
	nb.props = &props
	return nb
}

func (nb *NodeBuilder) Key(prop string) *AccessKey {
	return newAccessKey(nb, prop)
}

func (nb *NodeBuilder) buildLabels() string {
	var b strings.Builder
	template := ":$($%s)"

	for _, label := range nb.labels {
		ref := nb.nextTokenRef(label)
		nb.refs.token[label] = ref

		b.WriteString(fmt.Sprintf(template, ref))
	}
	return b.String()
}

func (nb *NodeBuilder) build() string {
	if nb.empty {
		return "()"
	}

	var b strings.Builder

	nb.refs.element = nb.nextNodeRef(nb)

	b.WriteByte('(')
	b.WriteString(nb.refs.element.String())
	b.WriteString(nb.buildLabels())

	if len(*nb.props) > 0 {
		nb.refs.props = nb.nextPropsRef(nb.props)
		b.WriteString(" $" + nb.refs.props.String())
	}

	b.WriteByte(')')

	return b.String()
}

func (nb *NodeBuilder) setBuilder(builder *Builder) {
	nb.Builder = builder
}

// for interface
func (*NodeBuilder) element() {}

type accessKeyPart interface {
	build(ref Ref) Ref
}
type funcPart struct {
	template string
}

func newFuncPart(template string) *funcPart {
	return &funcPart{template: template}
}

func (fp *funcPart) build(ref Ref) Ref {
	return Ref(fmt.Sprintf(fp.template, ref))
}

type AccessKey struct {
	owner     Element
	key       string
	funcParts accessKeyPart
}

func newAccessKey(owner Element, key string) *AccessKey {
	return &AccessKey{
		owner: owner,
		key:   key,
	}
}

func (ap *AccessKey) addPart(part accessKeyPart) {
	ap.funcParts = part
}

func (ap *AccessKey) ToLower() *AccessKey {
	ap.addPart(newFuncPart("toLower(%s)"))
	return ap
}

func (ap *AccessKey) ToUpper() *AccessKey {
	ap.addPart(newFuncPart("toUpper(%s)"))
	return ap
}

func (ap *AccessKey) toRef(builder *Builder) Ref {
	ap.owner.build()
	ref := ap.owner.GetRef()

	template := "%s[$%s]"
	fieldRef := builder.nextTokenRef(ap.key)

	newRef := Ref(fmt.Sprintf(template, ref, fieldRef))

	if ap.funcParts != nil {
		newRef = ap.funcParts.build(newRef)
	}

	return newRef
}
