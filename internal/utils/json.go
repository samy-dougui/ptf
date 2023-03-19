package utils

import "encoding/json"

func MarshalJson(content interface{}) ([]byte, error) {
	output, err := json.MarshalIndent(content, "", "  ")
	return output, err
}
