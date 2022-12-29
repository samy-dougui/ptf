package loader

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/ports"
)

func (l *Loader) LoadPlan(path string) (*ports.Plan, hcl.Diagnostics) {
	// TODO : better manage file loader
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
