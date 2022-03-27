package main

import (
	"fmt"
	"github.com/samy-dougui/tftest/cli"
	"os"
)

func main() {
	Execute()
}

func Execute() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
