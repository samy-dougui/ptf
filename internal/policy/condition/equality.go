package condition

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/utils"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
)

func Equality(attribute interface{}, expectedValue cty.Value) (bool, hcl.Diagnostic) {
	var isValid bool
	var diag = hcl.Diagnostic{}
	if expectedValue.Type().IsPrimitiveType() {
		isValid, diag = equalityPrimitive(&attribute, &expectedValue)
	} else if expectedValue.Type().IsObjectType() {
		var diags hcl.Diagnostics
		isValid, diags = equalityObject(&attribute, &expectedValue)
		if !isValid {
			diag.Detail = utils.ConcatDiagsDetail(&diags)
		}
	} else {
		log.Printf("Type un managed: %v", expectedValue.Type().FriendlyName())
	}
	return isValid, diag
}

func equalityPrimitive(attribute *interface{}, expectedValue *cty.Value) (bool, hcl.Diagnostic) {
	var isValid bool
	var diag hcl.Diagnostic
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(*expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = (*attribute).(float64) == expectedValueTyped
		if !isValid {
			diag.Detail = fmt.Sprintf("It was expecting %v, but it's equals to %v.", expectedValueTyped, (*attribute).(float64))
		}
	case cty.String:
		var expectedValueTyped string
		err := gocty.FromCtyValue(*expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = (*attribute).(string) == expectedValueTyped
		if !isValid {
			diag.Detail = fmt.Sprintf("It was expecting %v, but it's equals to %v.", expectedValueTyped, (*attribute).(string))
		}
	case cty.Bool:
		var expectedValueTyped bool
		err := gocty.FromCtyValue(*expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
		}
		isValid = (*attribute).(bool) == expectedValueTyped
		if !isValid {
			diag.Detail = fmt.Sprintf("It was expecting %v, but it's equals to %v.", expectedValueTyped, (*attribute).(bool))
		}
	default:
		log.Println(expectedValue.Type().IsPrimitiveType())
		diag.Detail = "Default Value"
		isValid = false
	}
	return isValid, diag
}

func equalityObject(attribute *interface{}, expectedValue *cty.Value) (bool, hcl.Diagnostics) {
	var diags hcl.Diagnostics
	attributeTyped := (*attribute).(map[string]interface{})
	for key, value := range expectedValue.AsValueMap() {
		if attributeValue, ok := attributeTyped[key]; ok {
			switch attributeValue.(type) {
			case string:
				if value.Type() != cty.String {
					diags = append(diags, &hcl.Diagnostic{
						Severity: hcl.DiagError,
						Detail:   fmt.Sprintf("Expected a %v for the key \"%v\", but got a string", value.Type().FriendlyName(), key),
					})
				} else {
					var expectedValueTyped string
					_ = gocty.FromCtyValue(value, &expectedValueTyped) // TODO: handle error
					if expectedValueTyped != attributeValue.(string) {
						diags = append(diags, &hcl.Diagnostic{
							Severity: hcl.DiagError,
							Detail:   fmt.Sprintf("The key \"%v\" was expecting %v but got %v", key, expectedValueTyped, attributeValue.(string)),
						})
					}
				}
			case float64:
				if value.Type() != cty.Number {
					diags = append(diags, &hcl.Diagnostic{
						Severity: hcl.DiagError,
						Detail:   fmt.Sprintf("Expected a %v for the key \"%v\", but got a float64", value.Type().FriendlyName(), key),
					})
				} else {
					var expectedValueTyped float64
					_ = gocty.FromCtyValue(value, &expectedValueTyped)
					if expectedValueTyped != attributeValue.(float64) {
						diags = append(diags, &hcl.Diagnostic{
							Severity: hcl.DiagError,
							Detail:   fmt.Sprintf("The key \"%v\" was expecting %v but got %v", key, expectedValueTyped, attributeValue.(float64)),
						})
					}
				}
			case bool:
				if value.Type() != cty.Bool {
					diags = append(diags, &hcl.Diagnostic{
						Severity: hcl.DiagError,
						Detail:   fmt.Sprintf("Expected a %v for the key \"%v\", but got a boolean", value.Type().FriendlyName(), key),
					})
				} else {
					var expectedValueTyped bool
					_ = gocty.FromCtyValue(value, &expectedValueTyped)
					if expectedValueTyped != attributeValue.(bool) {
						diags = append(diags, &hcl.Diagnostic{
							Severity: hcl.DiagError,
							Detail:   fmt.Sprintf("The key \"%v\" was expecting %v but got %v", key, expectedValueTyped, attributeValue.(bool)),
						})
					}
				}
			default:
				diags = append(diags, &hcl.Diagnostic{
					Severity: hcl.DiagError,
					Detail:   "The only permitted values for testing the equality of two dictionaries are string, float64 and boolean",
				})
			}
		} else {
			diags = append(diags, &hcl.Diagnostic{
				Severity: hcl.DiagError,
				Detail:   fmt.Sprintf("The key \"%v\" is not set", key),
			})
		}
	}
	return !diags.HasErrors(), diags
}
