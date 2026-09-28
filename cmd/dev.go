package main

import (
	"fmt"

	"github.com/deb84/neoqueries"
)

func main() {
	n := neoqueries.NewNodePattern()
	n2 := neoqueries.NewNodePattern()
	r := neoqueries.NewRelationPattern(neoqueries.RelationTo)
	pat := neoqueries.NewPattern().Node(n).Relation(r).Node(n2).Relation(r).Node(n2)
	nn := n.GetNode().GetRef().Prop("a")
	c := neoqueries.NewConditionBuilder().Prop(nn).Eq().Value("a")
	q, p := neoqueries.NewQueryBuilder().Match(pat).Where(c).Build()
	fmt.Println(q, p)

}
