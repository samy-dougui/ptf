package condition

import (
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
	"regexp"
)

// TODO: Recover from panic error handling
// TODO: Have a proper formatting for error

var OperatorMap = map[string]func(interface{}, cty.Value) bool{
	"=":  Equality,
	">":  SuperiorStrict,
	">=": SuperiorOrEqual,
	"<":  InferiorStrict,
	"<=": InferiorOrEqual,
	"re": RegexMatch,
}

func Equality(attribute interface{}, expectedValue cty.Value) bool {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		return attribute.(float64) == expectedValueTyped
	case cty.String:
		var expectedValueTyped string
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		return attribute.(string) == expectedValueTyped
	default:
		log.Println("Default value")
		return false
	}
}

func SuperiorStrict(attribute interface{}, expectedValue cty.Value) bool {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		return attribute.(float64) > expectedValueTyped
	default:
		log.Println("Only allowed type: number")
		return false
	}
}

func SuperiorOrEqual(attribute interface{}, expectedValue cty.Value) bool {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		return attribute.(float64) >= expectedValueTyped
	default:
		log.Println("Only allowed type: number")
		return false
	}
}

func InferiorStrict(attribute interface{}, expectedValue cty.Value) bool {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		return attribute.(float64) < expectedValueTyped
	default:
		log.Println("Only allowed type: number")
		return false
	}
}

func InferiorOrEqual(attribute interface{}, expectedValue cty.Value) bool {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		return attribute.(float64) <= expectedValueTyped
	default:
		log.Println("Only allowed type: number")
		return false
	}
}

func RegexMatch(attribute interface{}, expectedExpression cty.Value) bool {
	var expectedExpressionTyped string
	err := gocty.FromCtyValue(expectedExpression, &expectedExpressionTyped)
	if err != nil {
		log.Println(err)
	}
	match, _ := regexp.MatchString(expectedExpressionTyped, attribute.(string))
	return match
}
