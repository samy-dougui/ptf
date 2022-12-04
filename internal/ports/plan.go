package ports

type Plan struct {
	FormatVersion    string                 `json:"format_version"`
	TerraformVersion string                 `json:"terraform_version"`
	Variables        map[string]interface{} `json:"variables"`
	ResourceChanges  []ResourceChange       `json:"resource_changes"`
	Configuration    ProviderConfiguration  `json:"configuration"`
}

type ProviderConfiguration struct {
	Providers map[string]Provider `json:"provider_config"`
}

type ResourceChange struct {
	Address      string `json:"address"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	ProviderName string `json:"provider_name"`
	Index        string `json:"index,omitempty"`
	ActionReason string `json:"action_reason,omitempty"`
	Change       Change `json:"change"`
}

type Change struct {
	Actions         []string               `json:"actions"`
	Before          map[string]interface{} `json:"before"`
	After           map[string]interface{} `json:"after"`
	AfterUnknown    map[string]interface{} `json:"after_unknown"`
	BeforeSensitive map[string]interface{} `json:"before_sensitive"`
	AfterSensitive  map[string]interface{} `json:"after_sensitive"`
}
