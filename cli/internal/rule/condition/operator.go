package condition

import (
	"github.com/zclconf/go-cty/cty"
)

// TODO: find a way to have dynamic type in the configuration
// We want to have whatever type in the rule.hcl file
// Read the terraform plan, get the type of the attribute we are testing
// Try to convert the value we get from the rule.hcl file to the type we are getting from the json
// if we can't: raise an error saying the plan value does not have the correct type
// if we can: convert the value from the rule.hcl to the attribute file and execute the operator function

var OperatorMap = map[string]func(interface{}, cty.Value) bool{
	"=": Equality,
	//">":  Superior,
	//">=": SuperiorOrEqual,
	//"<":  Inferior,
	//"<=": InferiorOrEqual,
}

func Equality(attribute interface{}, expectedValue cty.Value) bool {
	//attributeType, _ := gocty.ToCtyValue(attribute, cty.String)
	//fmt.Println(attributeType.Type().)
	//fmt.Println(expectedValue.Type())
	//fmt.Println(expectedValue.Equals(attribute))
	//switch t := attribute.(type) {
	//case int:
	//	println("int")
	//	attributeTyped := attribute.(int)
	//	expectedValueTyped := expectedValue.AsString()
	//
	//case string:
	//	println("string")
	//	attributeTyped = attribute.(string)
	//default:
	//	log.Printf("unexpected type %v", t)
	//}
	return false
}

//func Superior(attribute interface{}, expectedValue cty.Value) bool {
//	return attribute.(int) > expectedValue.(int)
//}
//
//func SuperiorOrEqual(attribute interface{}, expectedValue cty.Value) bool {
//	return attribute.(int) >= expectedValue.(int)
//}
//
//func Inferior(attribute interface{}, expectedValue cty.Value) bool {
//	return attribute.(int) < expectedValue.(int)
//}
//
//func InferiorOrEqual(attribute interface{}, expectedValue cty.Value) bool {
//	return attribute.(int) <= expectedValue.(int)
//}
