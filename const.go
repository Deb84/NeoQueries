package neoqueries

var (
	DefaultNodeRef     Ref = Ref{"n"}
	DefaultRelationRef Ref = Ref{"r"}
	DefaultTokenRef    Ref = Ref{"t"}
	DefaultPropsRef    Ref = Ref{"p"}
	DefaultListRef     Ref = Ref{"ls"}
	DefaultValueRef    Ref = Ref{"v"}
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
