package neoqueries

type Element interface {
	buildElement(*BuildContext) Ref
	GetRef() *UnbuiltRef
}
