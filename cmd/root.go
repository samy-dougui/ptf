package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

const version string = "0.1.0-alpha"

var rootCmd = &cobra.Command{
	Use:   "ptf",
	Short: "ptf is a Policy as Code tool that lets you control your Terraform plan.",
	Long:  `ptf is a tool that helps you control your Terraform plan and state through configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
	Version: version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(ControlCmd)
	rootCmd.AddCommand(ServerCmd)
}
