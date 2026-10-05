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
