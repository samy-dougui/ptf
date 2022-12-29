package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testValidValues = []struct {
	NestedMap      interface{}
	AttributeQuery string
	Expected       interface{}
}{
	{
		NestedMap:      HASHMAP{"foo": "bar"},
		AttributeQuery: "foo",
		Expected:       []interface{}{"bar"},
	},
	{
		NestedMap: HASHMAP{
			"foo": HASHMAP{
				"bar": "hello",
			}},
		AttributeQuery: "foo.bar",
		Expected:       []interface{}{"hello"},
	},
	{
		NestedMap: HASHMAP{
			"foo": HASHMAP{
				"bar_0": "hello_0",
				"bar_1": "hello_1",
			},
			"bar": "foo",
		},
		AttributeQuery: "foo.bar_0",
		Expected:       []interface{}{"hello_0"},
	},
	{
		NestedMap: HASHMAP{
			"foo": []interface{}{
				HASHMAP{
					"bar": "hello_0",
				},
				HASHMAP{
					"bar": "hello_1",
				},
			},
		},
		AttributeQuery: "foo.[0].bar",
		Expected:       []interface{}{"hello_0"},
	},
	{
		NestedMap: HASHMAP{
			"foo": []interface{}{
				HASHMAP{
					"bar": "hello_0",
				},
				HASHMAP{
					"bar": "hello_1",
				},
			},
		},
		AttributeQuery: "foo.[1].bar",
		Expected:       []interface{}{"hello_1"},
	},
	{
		NestedMap: HASHMAP{
			"foo": []interface{}{
				HASHMAP{
					"bar": "hello_0",
				},
				HASHMAP{
					"bar": "hello_1",
				},
			},
		},
		AttributeQuery: "foo.[*].bar",
		Expected:       []interface{}{"hello_0", "hello_1"},
	},
	{
		NestedMap: HASHMAP{
			"foo": []interface{}{
				HASHMAP{
					"bar_0": "hello_0",
				},
				HASHMAP{
					"bar_1": HASHMAP{
						"hello": "world",
					},
				},
			},
		},
		AttributeQuery: "foo.[1].bar_1.hello",
		Expected:       []interface{}{"world"},
	},
	{
		NestedMap: []interface{}{
			HASHMAP{
				"foo": "bar_0",
			},
			HASHMAP{
				"foo": "bar_1",
			},
		},
		AttributeQuery: "[*].foo",
		Expected:       []interface{}{"bar_0", "bar_1"},
	},
	{
		NestedMap: []interface{}{
			HASHMAP{
				"foo": "bar_0",
			},
			HASHMAP{
				"foo": "bar_1",
			},
		},
		AttributeQuery: "[1].foo",
		Expected:       []interface{}{"bar_1"},
	},
	{
		NestedMap: HASHMAP{
			"foo": []interface{}{
				HASHMAP{
					"bar": "hello_0",
				},
				HASHMAP{
					"bar": "hello_1",
				},
			},
		},
		AttributeQuery: "foo",
		Expected: []interface{}{
			[]interface{}{
				HASHMAP{"bar": "hello_0"},
				HASHMAP{"bar": "hello_1"}},
		},
	},
}

func TestGetAttributeValid(t *testing.T) {
	for _, testValue := range testValidValues {
		attribute, err := GetResourceAttribute(testValue.NestedMap, testValue.AttributeQuery)
		assert.Nil(t, err)
		assert.Equal(t, testValue.Expected, attribute)
	}
}

func TestGetAttributeInvalid(t *testing.T) {
	assert.True(t, false)
}
