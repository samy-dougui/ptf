package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type HASHMAP = map[string]interface{}

var MissingAttributeError = errors.New("missing attribute")

//var AttributeNameFormatError = errors.New("improper attribute name format")

//var ValidAttributeNameRegex = regexp.MustCompile(`^[^\s\.]*[a-zA-Z0-9\*\.\[\]]+(?<!\.)(?<!\])$`) // TODO: handle this please

func GetResourceAttribute(nestedMap interface{}, attributeName string) ([]interface{}, error) {
	var attributes []interface{}
	var attributeNameParts = strings.Split(attributeName, ".")
	if len(attributeNameParts) == 1 {
		_attribute := nestedMap.(HASHMAP)[attributeNameParts[0]]
		if _attribute == nil {
			return nil, MissingAttributeError
		}
		return []interface{}{_attribute}, nil
	} else {
		if !strings.Contains(attributeNameParts[0], "[") {
			_attribute := nestedMap.(HASHMAP)
			nestedAttribute, err := GetResourceAttribute(_attribute[attributeNameParts[0]], strings.Join(attributeNameParts[1:], "."))
			if err != nil {
				return nil, err
			}
			attributes = append(attributes, nestedAttribute...)
		} else if attributeNameParts[0] != "[*]" {
			// We have something like "[integer]"
			var _attribute = nestedMap.([]interface{})
			listIndex, err := strconv.Atoi(strings.Trim(attributeNameParts[0], "[]"))
			if err != nil {
				// TODO: handle error
				fmt.Printf("When using the list indexing in the condition's attribute, the value between the [ ] needs to be an integer, here it's %v", listIndex)
			}
			if len(_attribute) < listIndex {
				// TODO: handle error
				fmt.Println("The value passed inside the [ ] is larger than the list, it has been replaced by the max value possible.")
				listIndex = len(_attribute) - 1
			}
			nestedAttributes, err := GetResourceAttribute(_attribute[listIndex], strings.Join(attributeNameParts[1:], "."))
			if err != nil {
				return nil, err
			}
			attributes = append(attributes, nestedAttributes...)
		} else {
			// We have "[*]"
			var _attributes = nestedMap.([]interface{})
			for _, _attribute := range _attributes {
				nestedAttributes, err := GetResourceAttribute(_attribute.(HASHMAP), strings.Join(attributeNameParts[1:], "."))
				if err != nil {
					return nil, err
				}
				attributes = append(attributes, nestedAttributes...)
			}
		}
	}
	return attributes, nil
}

//func ValidateAttributeName(attributeNameQuery string) bool {
//	return ValidAttributeNameRegex.MatchString(attributeNameQuery)
//}
