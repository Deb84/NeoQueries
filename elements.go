package neoqueries

type Element interface {
	buildElement(*Builder) Ref
	GetRef() *UnbuiltRef
}
