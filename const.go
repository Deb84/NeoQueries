package neoqueries

const (
	NodeRef     Ref = "n"
	RelationRef Ref = "r"
	TokenRef    Ref = "t"
	PropsRef    Ref = "p"
	ListRef     Ref = "ls"
	ValueRef    Ref = "v"
)

type ComparisonOperator string

const (
	Equal              ComparisonOperator = "="
	NotEqual           ComparisonOperator = "<>"
	LessThan           ComparisonOperator = "<"
	GreaterThan        ComparisonOperator = ">"
	LessThanOrEqual    ComparisonOperator = "<="
	GreaterThanOrEqual ComparisonOperator = ">="
)

type RelationDirectionString struct {
	Left  string
	Right string
}

var RelationTo = RelationDirectionString{
	Left:  "-",
	Right: "->",
}

var RelationUndirected = RelationDirectionString{
	Left:  "-",
	Right: "-",
}

type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}
