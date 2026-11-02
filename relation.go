package neoqueries

type Relation struct {
	ref Ref
}

func NewRelation() *Relation {
	return &Relation{}
}

func (r *Relation) GetRef() *UnbuiltRef {
	return newUnbuiltRef(r)
}

func (r *Relation) buildElement(context *BuildContext) Ref {
	r.ref = context.ensureRelationRef(r)
	return r.ref
}
