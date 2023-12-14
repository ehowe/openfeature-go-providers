package file_provider_test

import (
	"context"
	"reflect"

	"github.com/ehowe/openfeature-go-providers/providers/file_provider"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Flag Examples",
	func(file_name string, resolveType string, format string, flag_name string, defaultValue interface{}, expectedOutput interface{}) {
		provider := file_provider.NewFileProvider(file_name, format)
		var res any

		switch resolveType {
		case "boolean":
			res = provider.BooleanEvaluation(context.Background(), flag_name, defaultValue.(bool), map[string]interface{}{})
		case "string":
			res = provider.StringEvaluation(context.Background(), flag_name, defaultValue.(string), map[string]interface{}{})
		case "number":
			res = provider.IntEvaluation(context.Background(), flag_name, defaultValue.(int64), map[string]interface{}{})
		case "float":
			res = provider.FloatEvaluation(context.Background(), flag_name, defaultValue.(float64), map[string]interface{}{})
		}

		v := reflect.ValueOf(res)
		Expect(v.FieldByName("Value").Interface()).To(Equal(expectedOutput))
	},

	Entry("parses a boolean value", "./testdata/flags.json", "boolean", "json", "example_boolean", false, true),
	Entry("parses an invalid boolean value", "./testdata/flags.json", "boolean", "json", "example_invalid_boolean", false, false),
	Entry("parses a string value", "./testdata/flags.json", "string", "json", "example_named_string", "default", "medium"),
	Entry("parses an invalid string value", "./testdata/flags.json", "string", "json", "example_invalid_named_string", "default", "default"),
	Entry("parses a number value", "./testdata/flags.json", "number", "json", "example_named_number", int64(1), int64(128)),
	Entry("parses an invalid number value", "./testdata/flags.json", "number", "json", "example_invalid_named_number", int64(1), int64(1)),
	Entry("parses a float value", "./testdata/flags.json", "float", "json", "example_named_float", 1.0, 2.718281828459045),
	Entry("parses an invalid float value", "./testdata/flags.json", "float", "json", "example_invalid_named_float", 1.0, 1.0),
)
