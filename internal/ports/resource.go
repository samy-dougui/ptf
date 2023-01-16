package ports

type Resource struct {
	Address  string
	Type     string
	Provider string
	Values   map[string]interface{}
	Action   string
}

type InvalidResource struct {
	Address           string      `json:"address"`
	AttributeName     string      `json:"attribute_name"`
	ExpectedAttribute interface{} `json:"expected_attribute"`
	ReceivedAttribute interface{} `json:"received_attribute"`
	ErrorMessage      string      `json:"error_message"`
}

type NewInvalidResource struct {
	Address           string
	AttributeName     string
	ErrorMessage      string
	InvalidAttributes []InvalidAttribute
}

type InvalidAttribute struct {
	ExpectedValue interface{}
	ReceivedValue interface{}
}
