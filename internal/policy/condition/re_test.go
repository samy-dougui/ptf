package condition

import (
	"github.com/samy-dougui/ptf/internal/ports"
	"github.com/stretchr/testify/assert"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
	"testing"
)

func TestRegexMatchValid(t *testing.T) {
	input := "test_variable"
	regex := "[a-zA-Z]+_[a-zA-Z]+"
	regexInput, _ := gocty.ToCtyValue(regex, cty.String)
	isValid, invalidAttribute, err := RegexMatch(input, regexInput)
	assert.Nil(t, err)
	assert.True(t, isValid)
	assert.Equal(t, ports.InvalidAttribute{}, invalidAttribute)
}

func TestRegexMatchInvalid(t *testing.T) {
	input := "test_variable"
	regex := "[a-zA-Z]+-[a-zA-Z]+"
	regexInput, _ := gocty.ToCtyValue(regex, cty.String)
	isValid, invalidAttribute, err := RegexMatch(input, regexInput)
	assert.Nil(t, err)
	assert.False(t, isValid)
	assert.Equal(t, ports.InvalidAttribute{
		ReceivedValue: input,
		ExpectedValue: regex,
	}, invalidAttribute)
}
