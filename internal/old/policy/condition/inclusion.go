package condition

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/logging"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

func Inclusion(attribute interface{}, expectedValue cty.Value) (bool, hcl.Diagnostic) {
	logger := logging.GetLogger()
	var diag hcl.Diagnostic
	if !expectedValue.Type().IsTupleType() {
		// TODO: should implement a validation of the policy
		diag.Severity = hcl.DiagError
		diag.Detail = fmt.Sprintf("For operator \"in\": Expected a list, received a %v", expectedValue.Type().FriendlyName())
	} else {
		if allSameType := checkAllSameType(&expectedValue); !allSameType {
			logger.Debug("Not all the elements of the list have the same type.")
			diag.Detail = "We couldn't apply this policy as not all the elements of the provided list have the same type."
			return false, diag
		}

		tupleType := getType(&expectedValue)
		if !tupleType.IsPrimitiveType() {
			logger.Debugf("The allowed types in a list are number, string or boolean. This list has %v", tupleType.FriendlyName())
			diag.Detail = fmt.Sprintf("We couldn't apply this policy as the only types of element allowed in a list of value are number, string and boolean.")
			return false, diag
		}

		it := expectedValue.ElementIterator()
		attributeTyped, _ := gocty.ToCtyValue(attribute, tupleType)
		for it.Next() {
			if _, value := it.Element(); value.Equals(attributeTyped).True() {
				return true, diag
			}
		}
		diag.Detail = fmt.Sprintf("\"%v\" is not included in the provided list.", attribute)
		return false, diag
	}
	return false, diag
}

func NotInclusion(attribute interface{}, expectedValue cty.Value) (bool, hcl.Diagnostic) {
	logger := logging.GetLogger()
	var diag hcl.Diagnostic
	if !expectedValue.Type().IsTupleType() {
		// TODO: should implement a validation of the policy
		diag.Detail = fmt.Sprintf("For operator \"in\": Expected a list, received a %v", expectedValue.Type().FriendlyName())
		return false, diag
	} else {
		// TODO: Maybe put this in a validator class
		if allSameType := checkAllSameType(&expectedValue); !allSameType {
			logger.Debug("Not all the elements of the list have the same type.")
			diag.Detail = "We couldn't apply this policy as not all the elements of the provided list have the same type."
			return false, diag
		}

		tupleType := getType(&expectedValue)
		if !tupleType.IsPrimitiveType() {
			logger.Debugf("The allowed types in a list are number, string or boolean. This list has %v", tupleType.FriendlyName())
			diag.Detail = fmt.Sprintf("We couldn't apply this policy as the only types of element allowed in a list of value are number, string and boolean.")
			return false, diag
		}

		it := expectedValue.ElementIterator()
		attributeTyped, _ := gocty.ToCtyValue(attribute, tupleType)
		for it.Next() {
			if _, value := it.Element(); value.Equals(attributeTyped).True() {
				diag.Detail = fmt.Sprintf("\"%v\" is equals to one of the element of the provided list.", attribute)
				return false, diag
			}
		}
		return true, diag
	}
}

func checkAllSameType(expectedValue *cty.Value) bool {
	it := expectedValue.ElementIterator()
	expectedType := getType(expectedValue)
	for it.Next() {
		_, value := it.Element()
		if !value.Type().Equals(expectedType) {
			return false
		}
	}
	return true
}

func getType(expectedValue *cty.Value) cty.Type {
	return expectedValue.Type().TupleElementTypes()[0]
}
