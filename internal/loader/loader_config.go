package loader

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"strings"
)

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
