package loader

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"strings"
)

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
	Actions         []string               `json:"actions"`
	Before          map[string]interface{} `json:"before"`
	After           map[string]interface{} `json:"after"`
	AfterUnknown    interface{}            `json:"after_unknown"`
	BeforeSensitive interface{}            `json:"before_sensitive"`
	AfterSensitive  interface{}            `json:"after_sensitive"`
}

func (l *Loader) LoadPlan(path string) (*Plan, hcl.Diagnostics) {
	if exists, _ := l.FileSystem.Exists(path); !exists {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "File doesn't exist",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}
	jsonFileBytes, err := l.FileSystem.ReadFile(path)
	if err != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}

	var plan Plan
	_ = json.Unmarshal(jsonFileBytes, &plan)

	return &plan, nil
}

func (r *ResourceChange) GetAttribute(attribute string) interface{} {
	// TODO: it's only for the After, should be configurable
	var _attribute = r.Change.After
	var nestedAttributes = strings.Split(attribute, ".")
	for _, nestedAttribute := range nestedAttributes[:len(nestedAttributes)-1] {
		_attribute = _attribute[nestedAttribute].(map[string]interface{})
	}
	return _attribute[nestedAttributes[len(nestedAttributes)-1]]
}
