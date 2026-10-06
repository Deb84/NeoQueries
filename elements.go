package neoqueries

type ElementInterface interface {
	buildElement(*Builder) Ref
	GetRef() *UnbuiltRef
}
