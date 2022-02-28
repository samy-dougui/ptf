package loader

// Plan "mode" can be "managed", for resources, or "data", for data resources
type Plan struct {
	FormatVersion    string           `json:"format_version"`
	TerraformVersion string           `json:"terraform_version"`
	Variables        interface{}      `json:"variables"`
	PlannedValues    interface{}      `json:"planned_values"`
	ResourceChanges  []ResourceChange `json:"resource_changes"`
	PriorState       interface{}      `json:"prior_state"`
	Configuration    interface{}      `json:"configuration"`
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
	Actions         []string    `json:"actions"`
	Before          interface{} `json:"before"`
	After           interface{} `json:"after"`
	AfterUnknown    interface{} `json:"after_unknown"`
	BeforeSensitive interface{} `json:"before_sensitive"`
	AfterSensitive  interface{} `json:"after_sensitive"`
}
