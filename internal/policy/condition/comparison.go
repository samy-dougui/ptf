package condition

import (
	"fmt"
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
)

func SuperiorStrict(attribute interface{}, expectedValue cty.Value) (bool, ports.InvalidAttribute, error) {

	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
			return false, ports.InvalidAttribute{}, err
		}
		isValid := attribute.(float64) > expectedValueTyped
		if !isValid {
			return isValid, ports.InvalidAttribute{
				ExpectedValue: expectedValueTyped,
				ReceivedValue: attribute.(float64),
			}, nil
		} else {
			return true, ports.InvalidAttribute{}, nil
		}
	default:
		return false, ports.InvalidAttribute{}, fmt.Errorf("only allowed type for '>' operator is float64, received %s", expectedValue.Type().FriendlyName())

	}
}

func SuperiorOrEqual(attribute interface{}, expectedValue cty.Value) (bool, ports.InvalidAttribute, error) {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
			return false, ports.InvalidAttribute{}, err
		}
		isValid := attribute.(float64) >= expectedValueTyped
		if !isValid {
			return isValid, ports.InvalidAttribute{
				ExpectedValue: expectedValueTyped,
				ReceivedValue: attribute.(float64),
			}, nil
		} else {
			return true, ports.InvalidAttribute{}, nil
		}
	default:
		return false, ports.InvalidAttribute{}, fmt.Errorf("only allowed type for '>=' operator is float64, received %s", expectedValue.Type().FriendlyName())
	}
}

func InferiorStrict(attribute interface{}, expectedValue cty.Value) (bool, ports.InvalidAttribute, error) {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
			return false, ports.InvalidAttribute{}, err
		}
		isValid := attribute.(float64) < expectedValueTyped
		if !isValid {
			return isValid, ports.InvalidAttribute{
				ExpectedValue: expectedValueTyped,
				ReceivedValue: attribute.(float64),
			}, nil
		} else {
			return true, ports.InvalidAttribute{}, nil
		}
	default:
		return false, ports.InvalidAttribute{}, fmt.Errorf("only allowed type for '<' operator is float64, received %s", expectedValue.Type().FriendlyName())
	}
}

func InferiorOrEqual(attribute interface{}, expectedValue cty.Value) (bool, ports.InvalidAttribute, error) {
	switch expectedValue.Type() {
	case cty.Number:
		var expectedValueTyped float64
		err := gocty.FromCtyValue(expectedValue, &expectedValueTyped)
		if err != nil {
			log.Println(err)
			return false, ports.InvalidAttribute{}, err
		}
		isValid := attribute.(float64) <= expectedValueTyped
		if !isValid {
			return isValid, ports.InvalidAttribute{
				ExpectedValue: expectedValueTyped,
				ReceivedValue: attribute.(float64),
			}, nil
		} else {
			return true, ports.InvalidAttribute{}, nil
		}
	default:
		return false, ports.InvalidAttribute{}, fmt.Errorf("only allowed type for '<=' operator is float64, received %s", expectedValue.Type().FriendlyName())
	}
}
