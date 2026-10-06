package neoqueries

type Relation struct {
	ref Ref
}

func NewRelation() *Relation {
	return &Relation{}
}

func (r *Relation) getRef() Ref {
	return r.ref
}

func (r *Relation) GetRef() *UnbuiltRef {
	return newUnbuiltRef(r)
}

func (r *Relation) buildElement(builder *Builder) Ref {
	r.ref = builder.nextRelationRef(r)
	return r.ref
}
