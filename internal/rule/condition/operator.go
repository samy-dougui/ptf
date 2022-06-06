package condition

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
	"regexp"
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
}

func Equality(attribute interface{}, expectedValue cty.Value) (bool, hcl.Diagnostic) {
	var isValid bool
	var diag = hcl.Diagnostic{}
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = attribute.(float64) == expectedValueTyped
		diag.Detail = fmt.Sprintf("It was expecting %v, but it's equals to %v.", expectedValueTyped, attribute.(float64))
	case cty.String:
		var expectedValueTyped string
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = attribute.(string) == expectedValueTyped
		diag.Detail = fmt.Sprintf("It was expecting %v, but it's equals to %v.", expectedValueTyped, attribute.(string))
	default:
		diag.Detail = "Default Value"
		isValid = false
	}
	return isValid, diag
}

func SuperiorStrict(attribute interface{}, expectedValue cty.Value) (bool, hcl.Diagnostic) {
	var isValid bool
	var diag = hcl.Diagnostic{}
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = attribute.(float64) > expectedValueTyped
		diag.Detail = fmt.Sprintf("It was expecting to have a value stricly superior at %v, but it's equal to %v.", expectedValueTyped, attribute.(float64))
	default:
		diag.Detail = "Only allowed type: number"
		isValid = false
	}
	return isValid, diag
}

func SuperiorOrEqual(attribute interface{}, expectedValue cty.Value) (bool, hcl.Diagnostic) {
	var isValid bool
	var diag = hcl.Diagnostic{}
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = attribute.(float64) >= expectedValueTyped
		diag.Detail = fmt.Sprintf("It was expecting to have a value superior at %v, but it's equal to %v.", expectedValueTyped, attribute.(float64))
	default:
		diag.Detail = "Only allowed type: number"
		isValid = false
	}
	return isValid, diag
}

func InferiorStrict(attribute interface{}, expectedValue cty.Value) (bool, hcl.Diagnostic) {
	var isValid bool
	var diag = hcl.Diagnostic{}
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = attribute.(float64) < expectedValueTyped
		diag.Detail = fmt.Sprintf("It was expecting to have a value strictly inferior at %v, but it's equal to %v.", expectedValueTyped, attribute.(float64))
	default:
		diag.Detail = "Only allowed type: number"
		isValid = false
	}
	return isValid, diag
}

func InferiorOrEqual(attribute interface{}, expectedValue cty.Value) (bool, hcl.Diagnostic) {
	var isValid bool
	var diag = hcl.Diagnostic{}
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = attribute.(float64) <= expectedValueTyped
		diag.Detail = fmt.Sprintf("It was expecting to have a value inferior at %v, but it's equal to %v.", expectedValueTyped, attribute.(float64))
	default:
		diag.Detail = "Only allowed type: number"
		isValid = false
	}
	return isValid, diag
}

func RegexMatch(attribute interface{}, expectedExpression cty.Value) (bool, hcl.Diagnostic) {
	var diag = hcl.Diagnostic{}
	var expectedExpressionTyped string
	err := gocty.FromCtyValue(expectedExpression, &expectedExpressionTyped)
	if err != nil {
		log.Println(err)
	}
	isValid, _ := regexp.MatchString(expectedExpressionTyped, attribute.(string))
	if !isValid {
		diag.Detail = fmt.Sprintf("It was expecting to follow the pattern %v, but it's %v.", expectedExpressionTyped, attribute.(string))
	}
	return isValid, diag
}
