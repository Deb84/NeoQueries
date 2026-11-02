package neoqueries

import "slices"

import "errors"

type Builder struct {
	config Config
}

func NewBuilder(config *Config) (*Builder, error) {
	var c Config
	if config == nil {
		c = defaultConfig
	} else {
		c = *config
	}

	if err := checkConfig(c); err != nil {
		return nil, err
	}

	return &Builder{c}, nil
}

func checkConfig(config Config) error {
	if slices.Contains([]Ref{
		config.Ref.NodeRef,
		config.Ref.RelationRef,
		config.Ref.TokenRef,
		config.Ref.PropsRef,
		config.Ref.ListRef,
		config.Ref.ValueRef,
	}, "") {
		return errors.New("ref config can't be empty")
	}
	return nil
}

func (b *Builder) NewQueryBuilder() *QueryBuilder {
	return newQueryBuilderWithConfig(b.config)
}
