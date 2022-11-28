package ports

type Configuration struct {
	Variables              []Variable
	TerraformVersion       string
	ProvidersConfiguration []ProviderConfiguration
}

type Variable struct {
	Name  string
	Value interface{}
}

type ProviderConfiguration struct {
	Name              string
	VersionConstraint string
}
