package ports

type Configuration struct {
	Variables        []Variable
	TerraformVersion string
	Providers        []Provider
}

type Variable struct {
	Name  string
	Value interface{}
}

type Provider struct {
	Name              string                 `json:"name"`
	VersionConstraint string                 `json:"version_constraint"`
	Expressions       map[string]interface{} `json:"expressions"`
}
