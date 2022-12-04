package loader

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/logging"
	"github.com/samy-dougui/ptf/internal/ports"
	"strconv"
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

func (l *Loader) LoadPlan(path string) (*ports.Plan, hcl.Diagnostics) {
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

	var plan ports.Plan
	_ = json.Unmarshal(jsonFileBytes, &plan)

	return &plan, nil
}

// TODO: definitely need tests / to be moved to another package as it's not the plan loader jurisdiction
// TODO: need to update the rest to test for each returned attribute
// TODO: the naming of the variables should be revised

func GetAttributeNew(attribute interface{}, attributeName string) []interface{} {
	var splitAttribute = strings.Split(attributeName, ".")
	logger := logging.GetLogger()
	var attributes []interface{}
	if len(splitAttribute) == 1 {
		firstAttribute := splitAttribute[0]
		_attribute := attribute.(map[string]interface{})
		return []interface{}{_attribute[firstAttribute]}
	} else {
		firstAttribute := splitAttribute[0]
		if !strings.Contains(firstAttribute, "[") {
			var _attribute = attribute.(map[string]interface{})
			attributes = append(attributes, GetAttributeNew(_attribute[firstAttribute], strings.Join(splitAttribute[1:], "."))...)
		} else if firstAttribute != "[*]" {
			var _attribute = attribute.([]interface{})
			_listIndex := strings.Trim(firstAttribute, "[]")
			listIndex, err := strconv.Atoi(_listIndex)
			if err != nil {
				logger.Fatalf("When using the list indexing in the condition's attribute, the value between the [ ] needs to be an integer, here it's %v", _listIndex)
			}

			if len(_attribute) < listIndex {
				logger.Warn("The value passed inside the [ ] is larger than the list, it has been replaced by the max value possible.")
				listIndex = len(_attribute) - 1
			}
			attributes = append(attributes, GetAttributeNew(_attribute[listIndex], strings.Join(splitAttribute[1:], "."))...)
		} else {
			var _attributes = attribute.([]interface{})
			for _, _attribute := range _attributes {
				attributes = append(attributes, GetAttributeNew(_attribute.(map[string]interface{}), strings.Join(splitAttribute[1:], "."))...)
			}
		}
	}
	return attributes
}
