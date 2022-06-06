package main

import (
	"github.com/samy-dougui/ptf/cmd"
	"github.com/samy-dougui/ptf/internal/logging"
)

func main() {
	setUp()
	cmd.Execute()
}

func setUp() {
	logging.SetUpLogger()
	logging.SetUpDiagnosticLogger()
}
