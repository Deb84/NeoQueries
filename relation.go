package neoqueries

type Relation struct {
	ref RelationRef
}

func NewRelation() *Relation {
	return &Relation{}
}

func (r *Relation) GetRef() *UnbuiltRef[RelationRef] {
	return newUnbuiltRef(r)
}

func (r *Relation) buildElement(context *BuildContext) RelationRef {
	return context.ensureRelationRef(r)
}
