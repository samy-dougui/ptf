package cmd

import (
	"github.com/samy-dougui/ptf/internal/core"
	"github.com/samy-dougui/ptf/internal/loader"
	"github.com/samy-dougui/ptf/internal/policy"
	"github.com/samy-dougui/ptf/internal/utils"
	"github.com/samy-dougui/ptf/internal/ux"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	planPath        string
	policiesDirPath string
	prettyPrint     bool
	shortOutput     bool
	failOnWarning   bool
)

var ControlCmd = &cobra.Command{
	Use:   "control",
	Short: "Control your Terraform plan.",
	Long:  ``, // TODO: Add long description to control command

	Run: func(cmd *cobra.Command, args []string) {
		run(planPath, policiesDirPath)
	},
}

func init() {
	ControlCmd.Flags().StringVarP(&planPath, "plan", "p", "", "Path to the Terraform plan that needs to be controlled.")
	ControlCmd.Flags().StringVarP(&policiesDirPath, "chdir", "", ".", "Directory where the policy files are defined.")
	ControlCmd.Flags().BoolVarP(&prettyPrint, "pretty", "", true, "Print output of checks in a table.")
	ControlCmd.Flags().BoolVarP(&shortOutput, "short", "", false, "The output is just the table (only work in pretty mode).")
	ControlCmd.Flags().BoolVarP(&failOnWarning, "fail-on-warning", "", false, "Fail if there is at least one warning.")
	ControlCmd.MarkFlagRequired("plan")
}

func run(planPath string, policiesDirPath string) {
	normalizePoliciesDir, err := utils.NormalizePath(policiesDirPath)
	if err != nil {
		panic(err)
	}
	normalizePlanPath, err := utils.NormalizePath(planPath)
	if err != nil {
		panic(err)
	}
	policies, err := loader.LoadPolicies(normalizePoliciesDir)
	if err != nil {
		log.Printf("Error while loading directory %s: %s", policiesDirPath, err)
		return
	}
	resources, configuration := loader.LoadLocalResources(normalizePlanPath)
	validationOutput := core.Validate(&policies, &resources, &configuration)
	ux.Display(&validationOutput, prettyPrint, shortOutput)
	handleExit(&validationOutput, failOnWarning)
}

func handleExit(outputs *[]policy.Output, warningFailure bool) {
	warningCount := 0
	errorCount := 0
	for _, output := range *outputs {
		if !output.Validated {
			if output.Severity == policy.WARNING {
				warningCount += 1
			} else {
				errorCount += 1
			}
		}
	}
	if errorCount > 0 {
		os.Exit(1)
	} else {
		if warningCount > 0 && warningFailure {
			os.Exit(1)
		}
		os.Exit(0)
	}

}
