package neoqueries

import (
	"fmt"
	"strings"
)

type ValidRelationPattern interface {
	*RelationPattern | *RelationRefPattern
	build(context *BuildContext) PatternString
}

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

func (p *RelationPattern) GetDirection() RelationDirectionString {
	return p.direction
}

func (p *RelationPattern) Type(token string) *RelationPattern {
	p.token = token
	return p
}

func (p *RelationPattern) Props(props *Props) *RelationPattern {
	p.props = props
	return p
}

func (p *RelationPattern) build(context *BuildContext) PatternString {
	tokenTemplate := ":$($%s)"

	relationRef := p.relation.buildElement(context)

	var b strings.Builder

	b.WriteString(relationRef.String())

	if p.token != "" {
		tokenRef := context.ensureTokenRef(p.token)
		p.tokenRef[p.token] = tokenRef
		b.WriteString(fmt.Sprintf(tokenTemplate, tokenRef))
	}

	if len(*p.props) > 0 {
		propsRef := context.ensurePropsRef(p.props)
		b.WriteString(" $" + propsRef.String())
	}

	return PatternString(b.String())
}

type RelationRefPattern struct {
	ref *UnbuiltRef[RelationRef]
}

func (p *RelationRefPattern) build(context *BuildContext) PatternString {
	return PatternString(p.ref.buildRef(context).String())
}

type RelationPatternPart[T ValidRelationPattern] struct {
	pattern   T
	direction RelationDirectionString
}

func (p *RelationPatternPart[T]) build(context *BuildContext) PatternString {
	return PatternString(p.direction.Left + "[" + p.pattern.build(context).string() + "]" + p.direction.Right)
}
