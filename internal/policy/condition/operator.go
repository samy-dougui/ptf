package condition

import (
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/zclconf/go-cty/cty"
)

var OperatorMap = map[string]func(interface{}, cty.Value) (bool, ports.InvalidAttribute, error){
	// TODO: Recover from panic error handling
	// TODO: Have a proper formatting for error
	// TODO: Delete hcl.diagnostic
	"=":      Equality,
	">":      SuperiorStrict,
	">=":     SuperiorOrEqual,
	"<":      InferiorStrict,
	"<=":     InferiorOrEqual,
	"re":     RegexMatch,
	"in":     Inclusion,
	"not in": NotInclusion,
}
