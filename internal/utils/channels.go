package utils

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/samy-dougui/ptf/internal/logging"
	"sync"
)

// CloseChannel waits for the waiting group to be done and then closes the channel
func CloseChannel(wg *sync.WaitGroup, channel *chan hcl.Diagnostics) {
	wg.Wait()
	close(*channel)
}

func GatherDiagFromChannel(diagChannel *chan hcl.Diagnostics) hcl.Diagnostics {
	var diags hcl.Diagnostics

	for diag := range *diagChannel {
		_ = logging.WriteDiagnostics(diag)
		diags = append(diags, diag...)
	}
	return diags
}
