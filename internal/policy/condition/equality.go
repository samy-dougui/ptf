package condition

import (
	"errors"
	"fmt"
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
)

func Equality(attribute interface{}, expectedValue cty.Value) (bool, ports.InvalidAttribute, error) {
	if expectedValue.Type().IsPrimitiveType() {
		return equalityPrimitive(&attribute, &expectedValue)
	} else if expectedValue.Type().IsObjectType() {
		return equalityObject(&attribute, &expectedValue)
	} else {
		return false, ports.InvalidAttribute{}, fmt.Errorf("unsupported type %s", expectedValue.Type().FriendlyName())
	}
}

func equalityPrimitive(attribute *interface{}, expectedValue *cty.Value) (bool, ports.InvalidAttribute, error) {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(*expectedValue, &expectedValueTyped)
		if err != nil {
			return false, ports.InvalidAttribute{}, err
		}
		isValid := (*attribute).(float64) == expectedValueTyped
		if !isValid {
			return isValid, ports.InvalidAttribute{
				ExpectedValue: expectedValueTyped,
				ReceivedValue: (*attribute).(float64),
			}, nil
		} else {
			return true, ports.InvalidAttribute{}, nil
		}
	case cty.String:
		var expectedValueTyped string
		err := gocty.FromCtyValue(*expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid := (*attribute).(string) == expectedValueTyped
		if !isValid {
			return isValid, ports.InvalidAttribute{
				ExpectedValue: expectedValueTyped,
				ReceivedValue: (*attribute).(string),
			}, nil
		} else {
			return true, ports.InvalidAttribute{}, nil
		}
	case cty.Bool:
		var expectedValueTyped bool
		err := gocty.FromCtyValue(*expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid := (*attribute).(bool) == expectedValueTyped
		if !isValid {
			return isValid, ports.InvalidAttribute{
				ExpectedValue: expectedValueTyped,
				ReceivedValue: (*attribute).(bool),
			}, nil
		} else {
			return true, ports.InvalidAttribute{}, nil
		}
	default:
		return false, ports.InvalidAttribute{}, fmt.Errorf("invalid type %s", expectedValue.Type().FriendlyName())
	}
}

func equalityObject(attribute *interface{}, expectedValue *cty.Value) (bool, ports.InvalidAttribute, error) {
	return false, ports.InvalidAttribute{}, errors.New("dictionaries are not supported yet for the operator '='")
	//var diags hcl.Diagnostics
	//attributeTyped := (*attribute).(map[string]interface{})
	//for key, value := range expectedValue.AsValueMap() {
	//	if attributeValue, ok := attributeTyped[key]; ok {
	//		switch attributeValue.(type) {
	//		case string:
	//			if value.Type() != cty.String {
	//				diags = append(diags, &hcl.Diagnostic{
	//					Severity: hcl.DiagError,
	//					Detail:   fmt.Sprintf("Expected a %v for the key \"%v\", but got a string", value.Type().FriendlyName(), key),
	//				})
	//			} else {
	//				var expectedValueTyped string
	//				_ = gocty.FromCtyValue(value, &expectedValueTyped) // TODO: handle error
	//				if expectedValueTyped != attributeValue.(string) {
	//					diags = append(diags, &hcl.Diagnostic{
	//						Severity: hcl.DiagError,
	//						Detail:   fmt.Sprintf("The key \"%v\" was expecting %v but got %v", key, expectedValueTyped, attributeValue.(string)),
	//					})
	//				}
	//			}
	//		case float64:
	//			if value.Type() != cty.Number {
	//				diags = append(diags, &hcl.Diagnostic{
	//					Severity: hcl.DiagError,
	//					Detail:   fmt.Sprintf("Expected a %v for the key \"%v\", but got a float64", value.Type().FriendlyName(), key),
	//				})
	//			} else {
	//				var expectedValueTyped float64
	//				_ = gocty.FromCtyValue(value, &expectedValueTyped)
	//				if expectedValueTyped != attributeValue.(float64) {
	//					diags = append(diags, &hcl.Diagnostic{
	//						Severity: hcl.DiagError,
	//						Detail:   fmt.Sprintf("The key \"%v\" was expecting %v but got %v", key, expectedValueTyped, attributeValue.(float64)),
	//					})
	//				}
	//			}
	//		case bool:
	//			if value.Type() != cty.Bool {
	//				diags = append(diags, &hcl.Diagnostic{
	//					Severity: hcl.DiagError,
	//					Detail:   fmt.Sprintf("Expected a %v for the key \"%v\", but got a boolean", value.Type().FriendlyName(), key),
	//				})
	//			} else {
	//				var expectedValueTyped bool
	//				_ = gocty.FromCtyValue(value, &expectedValueTyped)
	//				if expectedValueTyped != attributeValue.(bool) {
	//					diags = append(diags, &hcl.Diagnostic{
	//						Severity: hcl.DiagError,
	//						Detail:   fmt.Sprintf("The key \"%v\" was expecting %v but got %v", key, expectedValueTyped, attributeValue.(bool)),
	//					})
	//				}
	//			}
	//		default:
	//			diags = append(diags, &hcl.Diagnostic{
	//				Severity: hcl.DiagError,
	//				Detail:   "The only permitted values for testing the equality of two dictionaries are string, float64 and boolean",
	//			})
	//		}
	//	} else {
	//		diags = append(diags, &hcl.Diagnostic{
	//			Severity: hcl.DiagError,
	//			Detail:   fmt.Sprintf("The key \"%v\" is not set", key),
	//		})
	//	}
	//}
	//return !diags.HasErrors(), diags
}
