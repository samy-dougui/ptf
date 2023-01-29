package ports

type Resource struct {
	Address  string
	Type     string
	Provider string
	Values   map[string]interface{}
	Action   string
}

type InvalidResource struct {
	Address           string             `json:"address"`
	AttributeName     string             `json:"attribute_name"`
	ErrorMessage      string             `json:"error_message"`
	InvalidAttributes []InvalidAttribute `json:"invalid_attributes"`
}

type InvalidAttribute struct {
	ExpectedValue interface{} `json:"expected_attribute,omitempty"`
	ReceivedValue interface{} `json:"received_attribute,omitempty"`
}
