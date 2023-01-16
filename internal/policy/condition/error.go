package condition

import "fmt"

type AttributeError struct {
	ExpectedAttribute interface{}
	ReceivedAttribute interface{}
}

func (a AttributeError) Error() string {
	return fmt.Sprintf("expecting the attribute %v, received %v", a.ExpectedAttribute, a.ReceivedAttribute)
}
