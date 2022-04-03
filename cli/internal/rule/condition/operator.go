package condition

var OperatorMap = map[string]func(interface{}, interface{}) bool{
	"=": Equality,
}

func Equality(attribute interface{}, expectedValue interface{}) bool {
	return attribute.(string) == expectedValue.(string)
}
