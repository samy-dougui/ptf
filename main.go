package main

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	loader2 "github.com/samy-dougui/tftest/internal/loader"
	"github.com/samy-dougui/tftest/internal/resource"
	"github.com/samy-dougui/tftest/internal/variable"
	"github.com/zclconf/go-cty/cty"
	"os"
)

func main() {
	wr := hcl.NewDiagnosticTextWriter(
		os.Stdout,                    // writer to send messages to
		hclparse.NewParser().Files(), // the parser's file cache, for source snippets
		78,                           // wrapping width
		true,                         // generate colored/highlighted output
	)
	var diags hcl.Diagnostics
	var loader loader2.Loader
	loader.Init()

	configValue, testConfigDiags := loader.LoadTestConfig("./example/config.yaml")
	diags = append(diags, testConfigDiags...)
	// TODO: LoadHCLFile should used the path given in the config file
	body, diagHCLFile := loader.LoadHCLFile("./example/data/module_1/main.tf")
	diags = append(diags, diagHCLFile...)

	content, _ := body.Content(configFileSchema)
	variables := map[string]*variable.Variable{}
	varDiags := ExtractVariables(content, &variables)
	diags = append(diags, varDiags...)

	err := AssignVariablesValue(&variables, configValue)
	if err != nil {
		fmt.Printf("Error assigning the values")
	}

	ctx := CreateContext(&variables)
	resourceDiags := ExtractRessources(content, &ctx)
	diags = append(diags, resourceDiags...)
	errDiag := wr.WriteDiagnostics(diags)
	if errDiag != nil {
		fmt.Printf("Error while writing the diagnostics: %v", errDiag)
	}
}

func CreateContext(variables *map[string]*variable.Variable) hcl.EvalContext {
	data := make(map[string]cty.Value)
	for varName, varValues := range *variables {
		data[varName] = varValues.Value
	}
	return hcl.EvalContext{
		Variables: map[string]cty.Value{
			"var": cty.ObjectVal(data),
			"local": cty.ObjectVal(map[string]cty.Value{
				"super_var_2": cty.StringVal("add_spice_2")}),
		},
	}
}
func ExtractVariables(content *hcl.BodyContent, variables *map[string]*variable.Variable) hcl.Diagnostics {
	var diags hcl.Diagnostics
	for _, block := range content.Blocks {
		switch block.Type {
		case "variable":
			newVariable := &variable.Variable{Name: block.Labels[0]}
			varBlockContent, _, variableDiags := block.Body.PartialContent(variable.BlockSchema)
			diags = append(diags, variableDiags...)
			if defaultValue, exists := varBlockContent.Attributes["default"]; exists {
				value, defaultValueDiags := defaultValue.Expr.Value(nil)
				diags = append(diags, defaultValueDiags...)
				newVariable.Default = value
			}
			if description, exists := varBlockContent.Attributes["description"]; exists {
				descriptionDiags := gohcl.DecodeExpression(description.Expr, nil, &newVariable.Description)
				diags = append(diags, descriptionDiags...)
			}
			(*variables)[newVariable.Name] = newVariable
		default:
			continue
		}
	}
	return diags
}
func ExtractRessources(content *hcl.BodyContent, ctx *hcl.EvalContext) hcl.Diagnostics {
	var diags hcl.Diagnostics
	for _, block := range content.Blocks {
		switch block.Type {
		case "resource":
			_, remain, resourceDiags := block.Body.PartialContent(resource.BlockSchema)
			diags = append(diags, resourceDiags...)
			attrs, attributesDiags := remain.JustAttributes()
			diags = append(diags, attributesDiags...)
			for name, attr := range attrs {
				value, _ := attr.Expr.Value(ctx)
				fmt.Printf("Resource Type: %v Resource name: %v, attribute name: %v, value from config: %v\n", block.Labels[0], block.Labels[1], name, value.AsString())
			}
		default:
			continue
		}
	}
	return diags
}

func AssignVariablesValue(variables *map[string]*variable.Variable, config map[string]string) error {
	for variableName, variableValue := range *variables {
		if variableConfigValue, exists := config[variableName]; exists {
			variableValue.Value = cty.StringVal(variableConfigValue)
		} else {
			if variableValue.Default != cty.NilVal {
				variableValue.Value = variableValue.Default
			}
		}
	}
	return nil
}

var configFileSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "terraform",
		},
		{
			// This one is not really valid, but we include it here so we
			// can create a specialized error message hinting the user to
			// nest it inside a "terraform" block.
			Type: "required_providers",
		},
		{
			Type:       "provider",
			LabelNames: []string{"name"},
		},
		{
			Type:       "variable",
			LabelNames: []string{"name"},
		},
		{
			Type: "locals",
		},
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
		{
			Type:       "resource",
			LabelNames: []string{"type", "name"},
		},
		{
			Type:       "data",
			LabelNames: []string{"type", "name"},
		},
		{
			Type: "moved",
		},
	},
}
