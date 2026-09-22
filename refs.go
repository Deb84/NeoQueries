package queries

import "NeoQueries"

type QueryRef interface {
	toRef(b *NeoQueries.Builder) Ref
}

type Ref string

func (r Ref) String() string {
	return string(r)
}

func (r Ref) toRef(b *NeoQueries.Builder) Ref {
	return r
}
