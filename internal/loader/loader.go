package loader

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/spf13/afero"
	"strings"
)

type Loader struct {
	Parser     *hclparse.Parser
	FileSystem afero.Afero
}

func (l *Loader) Init() {
	l.Parser = hclparse.NewParser()
	l.FileSystem = afero.Afero{
		Fs: afero.NewOsFs(),
	}
}

func (l *Loader) LoadHCLFile(path string) (hcl.Body, hcl.Diagnostics) {
	src, err := l.FileSystem.ReadFile(path)
	if err != nil {
		return nil, hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read file",
				Detail:   fmt.Sprintf("The file %q could not be read.", path),
			},
		}
	}
	var file *hcl.File
	var diags hcl.Diagnostics
	switch {
	case strings.HasSuffix(path, ".json"):
		file, diags = l.Parser.ParseJSON(src, path)
	default:
		file, diags = l.Parser.ParseHCL(src, path)
	}
	if file == nil || file.Body == nil {
		return hcl.EmptyBody(), diags
	}

	return file.Body, diags
}

func (l *Loader) LoadPlan(path string) (*Plan, hcl.Diagnostics) {
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
