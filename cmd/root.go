package cmd

import (
	"github.com/samy-dougui/ptf/cmd/control"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ptf",
	Short: "ptf helps you control your Terraform plan and Terraform state",
	Long:  `ptf is a tool that helps you control your Terraform plan and state through config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(control.ControlCmd)
}
