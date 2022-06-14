package condition

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

// TODO: Recover from panic error handling
// TODO: Have a proper formatting for error

var OperatorMap = map[string]func(interface{}, cty.Value) (bool, hcl.Diagnostic){
	"=":  Equality,
	">":  SuperiorStrict,
	">=": SuperiorOrEqual,
	"<":  InferiorStrict,
	"<=": InferiorOrEqual,
	"re": RegexMatch,
	"in": Inclusion,
	"not in": NotInclusion,
}
