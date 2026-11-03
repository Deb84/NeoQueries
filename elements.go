package neoqueries

type Element[T Reference] interface {
	buildElement(*BuildContext) T
	GetRef() *UnbuiltRef[T]
}
