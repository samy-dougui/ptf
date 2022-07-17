package condition

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
)

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
