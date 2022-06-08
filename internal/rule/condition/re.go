package condition

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"log"
	"regexp"
)

func RegexMatch(attribute interface{}, expectedExpression cty.Value) (bool, hcl.Diagnostic) {
	var diag = hcl.Diagnostic{}
	var expectedExpressionTyped string
	err := gocty.FromCtyValue(expectedExpression, &expectedExpressionTyped)
	if err != nil {
		log.Println(err)
	}
	isValid, _ := regexp.MatchString(expectedExpressionTyped, attribute.(string))
	if !isValid {
		diag.Detail = fmt.Sprintf("It was expecting to follow the pattern %v, but the value is \"%v\".", expectedExpressionTyped, attribute.(string))
	}
	return isValid, diag
}
