package cmd

import (
	"github.com/samy-dougui/ptf/internal/core"
	"github.com/samy-dougui/ptf/internal/loader"
	"github.com/samy-dougui/ptf/internal/utils"
	"github.com/samy-dougui/ptf/internal/ux"
	"github.com/spf13/cobra"
	"log"
)

var (
	planPath        string
	policiesDirPath string
	prettyPrint     bool
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
	ux.PrettyDisplay(&validationOutput)
	// TODO: os exit if failure => flag fail on warning
	return
}
