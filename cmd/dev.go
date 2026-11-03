package main

import (
	"fmt"

	"github.com/deb84/neoqueries"
)

type Relation struct {
	Type  string
	Props *neoqueries.Props
	N1    *Node
	N2    *Node
}

type Node struct {
	Label string
	Props *neoqueries.Props
}

func newNodePattern(node *Node) *neoqueries.NodePattern {
	pattern := neoqueries.NewNodePattern()

	if node.Label != "" {
		pattern.Label(node.Label)
	}
	if node.Props != nil {
		pattern.Props(node.Props)
	}

	return pattern
}

func newRelationPattern(relation *Relation) *neoqueries.RelationPattern {
	pattern := neoqueries.NewRelationPattern(neoqueries.RelationTo)

	if relation.Type != "" {
		pattern.Type(relation.Type)
	}
	if relation.Props != nil {
		pattern.Props(relation.Props)
	}

	return pattern
}

func CreateRelationQuery(relation *Relation) *neoqueries.QueryBuilder {
	n1 := newNodePattern(relation.N1)
	n2 := newNodePattern(relation.N2)

	return neoqueries.NewQueryBuilder().
		Match(n1).
		Match(n2).
		Create(neoqueries.NewPattern().
			NodeRef(n1.GetNode().GetRef()).
			Relation(newRelationPattern(relation)).
			NodeRef(n2.GetNode().GetRef()))
}

func main() {
	fmt.Println(CreateRelationQuery(&Relation{
		Type: "relationType",
		Props: &neoqueries.Props{
			"a": "b",
		},
		N1: &Node{
			Label: "labelN1",
			Props: &neoqueries.Props{
				"propsN1k": "propsN1v",
			},
		},
		N2: &Node{
			Label: "labelN2",
			Props: &neoqueries.Props{
				"propsN2k": "propsN2v",
			},
		},
	}).Build())
}
