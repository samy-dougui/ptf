package condition

import (
	"fmt"
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
)

func Inclusion(attribute interface{}, expectedValue cty.Value) (bool, ports.InvalidAttribute, error) {
	if !expectedValue.Type().IsTupleType() {
		return false, ports.InvalidAttribute{}, fmt.Errorf("the 'in' operator should be used with a list, but received %s", expectedValue.Type().FriendlyName())
	} else {
		if allSameType := checkAllSameType(&expectedValue); !allSameType {
			log.Println("not all the same type")
			return false, ports.InvalidAttribute{}, fmt.Errorf("the 'in' operator only works with a list of the same type")
		}

		tupleType := getType(&expectedValue)
		if !tupleType.IsPrimitiveType() {
			return false, ports.InvalidAttribute{}, fmt.Errorf("the 'in' operator only work for a list of element of type number, string and boolean")
		}
		it := expectedValue.ElementIterator()
		attributeTyped, _ := gocty.ToCtyValue(attribute, tupleType)
		for it.Next() {
			if _, value := it.Element(); value.Equals(attributeTyped).True() {
				return true, ports.InvalidAttribute{}, nil
			}
		}
		return false, ports.InvalidAttribute{
			ReceivedValue: attribute,
		}, nil
	}
}

func NotInclusion(attribute interface{}, expectedValue cty.Value) (bool, ports.InvalidAttribute, error) {
	if !expectedValue.Type().IsTupleType() {
		// TODO: should implement a validation of the policy
		return false, ports.InvalidAttribute{}, fmt.Errorf("the 'not in' operator should be used with a list, but received %s", expectedValue.Type().FriendlyName())
	} else {
		// TODO: Maybe put this in a validator class
		if allSameType := checkAllSameType(&expectedValue); !allSameType {
			return false, ports.InvalidAttribute{}, fmt.Errorf("the 'not in' operator only works with a list of the same type")
		}

		tupleType := getType(&expectedValue)
		if !tupleType.IsPrimitiveType() {
			return false, ports.InvalidAttribute{}, fmt.Errorf("the 'not in' operator only work for a list of element of type number, string and boolean")
		}

		it := expectedValue.ElementIterator()
		attributeTyped, _ := gocty.ToCtyValue(attribute, tupleType)
		for it.Next() {
			if _, value := it.Element(); value.Equals(attributeTyped).True() {
				return false, ports.InvalidAttribute{
					ReceivedValue: attribute,
				}, nil
			}
		}
		return true, ports.InvalidAttribute{}, nil
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
