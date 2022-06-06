package main

import (
	"fmt"
	"github.com/samy-dougui/tftest/cli"
	"github.com/samy-dougui/tftest/cli/logging"
	"os"
)

func main() {
	setUp()
	logger := logging.GetLogger()
	logger.Info("test info")
	//Execute()
}

func Execute() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setUp() {
	logging.SetUpLogger()
}
