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
	NodeRef:     NodeRef,
	RelationRef: RelationRef,
	TokenRef:    TokenRef,
	PropsRef:    PropsRef,
	ListRef:     ListRef,
	ValueRef:    ValueRef,
}

var defaultConfig = Config{
	Ref: defaultRefConfig,
}
