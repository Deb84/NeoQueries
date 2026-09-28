package neoqueries

import (
	"fmt"
	"strings"
)

type RelationPattern struct {
	*patternPart
	relation  *Relation
	token     string
	direction RelationDirectionString
}

func NewRelationPattern(direction RelationDirectionString) *RelationPattern {
	return &RelationPattern{
		patternPart: newPatternPart(),
		relation:    NewRelation(),
		direction:   direction,
	}
}

func (p *RelationPattern) GetRelation() *Relation {
	return p.relation
}

func (p *RelationPattern) Type(token string) *RelationPattern {
	p.token = token
	return p
}

func (p *RelationPattern) Props(props *Props) *RelationPattern {
	p.props = props
	return p
}

func (p *RelationPattern) build(pattern *Pattern) PatternString {
	tokenTemplate := ":$($%s)"

	p.relation.setBuilder(pattern.builder)
	relationRef := p.relation.buildElement()

	var b strings.Builder

	b.WriteString(p.direction.Left)
	b.WriteByte('[')
	b.WriteString(relationRef.String())

	if p.token != "" {
		tokenRef := pattern.builder.nextTokenRef(p.token)
		p.tokenRef[p.token] = tokenRef
		b.WriteString(fmt.Sprintf(tokenTemplate, tokenRef))
	}

	if len(*p.props) > 0 {
		p.propsRef = pattern.builder.nextPropsRef(p.props)
		b.WriteString(" $" + p.propsRef.String())
	}

	b.WriteByte(']')
	b.WriteString(p.direction.Right)

	return PatternString(b.String())
}
