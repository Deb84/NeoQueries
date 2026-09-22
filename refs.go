package neoqueries

type QueryRef interface {
	toRef(*Builder) Ref
}

type Ref string

func (r Ref) String() string {
	return string(r)
}

func (r Ref) toRef(*Builder) Ref {
	return r
}
