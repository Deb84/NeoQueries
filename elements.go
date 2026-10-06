package neoqueries

type ElementInterface interface {
	buildElement(*Builder) Ref
	getRef() Ref
	GetRef() *UnbuiltRef
}
