package neoqueries

import (
	"fmt"
	"strings"
)

type NodePattern struct {
	*patternPart
	node        *Node
	labels      []string
	accessProps []string
	empty       bool
}

func NewNodePattern() *NodePattern {
	return &NodePattern{
		patternPart: newPatternPart(),
		node:        NewNode(),
	}
}

func (p *NodePattern) GetNode() *Node {
	return p.node
}

func (p *NodePattern) Empty() *NodePattern {
	p.empty = true
	return p
}

func (p *NodePattern) Label(label string) *NodePattern {
	p.labels = append(p.labels, label)
	return p
}

func (p *NodePattern) Props(props *Props) *NodePattern {
	p.props = props
	return p
}

func (p *NodePattern) buildLabels(builder *Builder) string {
	var b strings.Builder
	template := ":$($%s)"

	for _, label := range p.labels {
		ref := builder.nextTokenRef(label)
		p.tokenRef[label] = ref

		b.WriteString(fmt.Sprintf(template, ref))
	}
	return b.String()
}

func (p *NodePattern) build(builder *Builder) PatternString {
	if p.empty {
		return "()"
	}
	var b strings.Builder

	nodeRef := p.node.buildElement(builder)

	b.WriteByte('(')
	b.WriteString(nodeRef.String())
	b.WriteString(p.buildLabels(builder))

	if len(*p.props) > 0 {
		p.propsRef = builder.nextPropsRef(p.props)
		b.WriteString(" $" + p.propsRef.String())
	}

	b.WriteByte(')')

	return PatternString(b.String())
}
