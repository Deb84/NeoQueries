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
	r1 := neoqueries.NewRelationPattern(neoqueries.RelationTo)
	_ = r1

	p := neoqueries.NewPattern().
		Node(n1).
		Relation(r1).
		Node(n2)

	p2 := neoqueries.NewPattern().
		NodeRef(n1.GetNode().GetRef()).
		RelationRef(r1.GetRelation().GetRef(), r1.GetDirection()).
		NodeRef(n2.GetNode().GetRef())

	return neoqueries.NewQueryBuilder().
		Match(n1, n2, p2).
		Create(p)
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
