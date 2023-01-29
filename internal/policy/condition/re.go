package condition

import (
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
	"regexp"
)

func RegexMatch(attribute interface{}, expectedExpression cty.Value) (bool, ports.InvalidAttribute, error) {
	var expectedExpressionTyped string
	err := gocty.FromCtyValue(expectedExpression, &expectedExpressionTyped)
	if err != nil {
		log.Println(err)
		return false, ports.InvalidAttribute{}, err
	}
	isValid, _ := regexp.MatchString(expectedExpressionTyped, attribute.(string))
	if !isValid {
		return false, ports.InvalidAttribute{
			ReceivedValue: attribute.(string),
			ExpectedValue: expectedExpressionTyped,
		}, nil
	} else {
		return true, ports.InvalidAttribute{}, nil
	}
}
