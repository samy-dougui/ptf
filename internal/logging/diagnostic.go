package logging

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"os"
)

var diagnosticLogger hcl.DiagnosticWriter

func SetUpDiagnosticLogger() {
	diagnosticLogger = hcl.NewDiagnosticTextWriter(
		os.Stdout,
		hclparse.NewParser().Files(),
		78,
		true,
	)
}

func WriteDiagnostics(diagnostics hcl.Diagnostics) error {
	err := diagnosticLogger.WriteDiagnostics(diagnostics)
	return err
}
