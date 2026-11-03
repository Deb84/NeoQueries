package neoqueries

type Config struct {
	Ref RefConfig
}

type RefConfig struct {
	NodeRef     Ref
	RelationRef Ref
	TokenRef    Ref
	PropsRef    Ref
	ListRef     Ref
	ValueRef    Ref
}

var defaultRefConfig = RefConfig{
	NodeRef:     DefaultNodeRef,
	RelationRef: DefaultRelationRef,
	TokenRef:    DefaultTokenRef,
	PropsRef:    DefaultPropsRef,
	ListRef:     DefaultListRef,
	ValueRef:    DefaultValueRef,
}

var defaultConfig = Config{
	Ref: defaultRefConfig,
}
