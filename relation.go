package queries

import (
	"NeoQueries"
	"fmt"
	"strings"
)

type RelationNodes struct {
	N1 *NodeBuilder
	N2 *NodeBuilder
}

type RelationDirection struct {
	relationBuilder *RelationBuilder
}

func (rd *RelationDirection) setDirection(n1, n2 *NodeBuilder, rds RelationDirectionString) {
	rd.relationBuilder.Nodes.N1 = n1
	rd.relationBuilder.Nodes.N2 = n2
	rd.relationBuilder.direction = rds
}

func (rd *RelationDirection) To(from, to *NodeBuilder) *RelationBuilder {
	rd.setDirection(from, to, RelationTo)
	return rd.relationBuilder
}

func (rd *RelationDirection) Undirected(n1, n2 *NodeBuilder) *RelationBuilder {
	rd.setDirection(n1, n2, RelationUndirected)
	return rd.relationBuilder
}

type RelationBuilder struct {
	*ElementBuilder
	token     string
	direction RelationDirectionString
	Nodes     RelationNodes
}

func NewRelationBuilder() *RelationDirection {
	rb := &RelationBuilder{
		ElementBuilder: newElementBuilder(),
	}
	return &RelationDirection{
		relationBuilder: rb,
	}
}

func (rb *RelationBuilder) Type(token string) *RelationBuilder {
	rb.token = token
	return rb
}

func (rb *RelationBuilder) Props(props NeoQueries.Props) *RelationBuilder {
	rb.props = &props
	return rb
}

func (rb *RelationBuilder) build() string {
	tokenTemplate := ":$($%s)"

	rb.refs.element = rb.nextRelationRef(rb)

	var b strings.Builder

	n1Query := rb.Nodes.N1.build()
	n2Query := rb.Nodes.N2.build()

	b.WriteString(n1Query)
	b.WriteString(rb.direction.Left)
	b.WriteByte('[')
	b.WriteString(rb.refs.element.String())

	if rb.token != "" {
		tokenRef := rb.nextTokenRef(rb.token)
		rb.refs.token[rb.token] = tokenRef
		b.WriteString(fmt.Sprintf(tokenTemplate, tokenRef))
	}

	if len(*rb.props) > 0 {
		rb.refs.props = rb.nextPropsRef(rb.props)
		b.WriteString(" $" + rb.refs.props.String())
	}

	b.WriteByte(']')
	b.WriteString(rb.direction.Right)
	b.WriteString(n2Query)

	return b.String()
}

func (rb *RelationBuilder) setBuilder(builder *NeoQueries.Builder) {
	rb.Builder = builder
	rb.Nodes.N1.setBuilder(builder)
	rb.Nodes.N2.setBuilder(builder)
}

// for interface
func (*RelationBuilder) element() {}
