package neoqueries

type Relation struct {
	*Element
	ref Ref
}

func NewRelation() *Relation {
	return &Relation{
		Element: newElement(),
	}
}

func (r *Relation) getRef() Ref {
	return r.ref
}

func (r *Relation) GetRef() *UnbuiltRef {
	return newUnbuiltRef(r)
}

func (r *Relation) buildElement() Ref {
	r.ref = r.builder.nextRelationRef(r)
	return r.ref
}
