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

func (p *NodePattern) buildLabels(context *BuildContext) string {
	var b strings.Builder
	template := ":$($%s)"

	for _, label := range p.labels {
		ref := context.ensureTokenRef(label)
		p.tokenRef[label] = ref

		b.WriteString(fmt.Sprintf(template, ref))
	}
	return b.String()
}

func (p *NodePattern) build(context *BuildContext) PatternNodeString {
	if p.empty {
		return PatternNodeString{"()"}
	}

	var b strings.Builder

	nodeRef := p.node.buildElement(context)

	b.WriteByte('(')
	b.WriteString(nodeRef.String())
	b.WriteString(p.buildLabels(context))

	if len(*p.props) > 0 {
		propsRef := context.ensurePropsRef(p.props)
		b.WriteString(" $" + propsRef.String())
	}

	b.WriteByte(')')

	patternString := PatternString(b.String())
	return PatternNodeString{patternString}
}
