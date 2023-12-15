package file_provider_test

import (
	"context"
	"reflect"

	"github.com/ehowe/openfeature-go-providers/providers/file_provider"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Flag Examples",
	func(file_name string, resolveType string, format string, flag_name string, defaultValue interface{}, expectedOutput interface{}, deepKeys ...string) {
		provider := file_provider.NewFileProvider(file_name, format, deepKeys)
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

	Entry("json parses a boolean value", "./testdata/flags.json", "boolean", "json", "example_boolean", false, true),
	Entry("json parses an invalid boolean value", "./testdata/flags.json", "boolean", "json", "example_invalid_boolean", false, false),
	Entry("json parses a string value", "./testdata/flags.json", "string", "json", "example_named_string", "default", "medium"),
	Entry("json parses an invalid string value", "./testdata/flags.json", "string", "json", "example_invalid_named_string", "default", "default"),
	Entry("json parses a number value", "./testdata/flags.json", "number", "json", "example_named_number", int64(1), int64(128)),
	Entry("json parses an invalid number value", "./testdata/flags.json", "number", "json", "example_invalid_named_number", int64(1), int64(1)),
	Entry("json parses a float value", "./testdata/flags.json", "float", "json", "example_named_float", 1.0, 2.718281828459045),
	Entry("json parses an invalid float value", "./testdata/flags.json", "float", "json", "example_invalid_named_float", 1.0, 1.0),

	Entry("json parses a boolean value", "./testdata/deeply_nested_flags.json", "boolean", "json", "example_boolean", false, true, "deeply", "nested"),
	Entry("json parses an invalid boolean value", "./testdata/deeply_nested_flags.json", "boolean", "json", "example_invalid_boolean", false, false, "deeply", "nested"),
	Entry("json parses a string value", "./testdata/deeply_nested_flags.json", "string", "json", "example_named_string", "default", "medium", "deeply", "nested"),
	Entry("json parses an invalid string value", "./testdata/deeply_nested_flags.json", "string", "json", "example_invalid_named_string", "default", "default", "deeply", "nested"),
	Entry("json parses a number value", "./testdata/deeply_nested_flags.json", "number", "json", "example_named_number", int64(1), int64(128), "deeply", "nested"),
	Entry("json parses an invalid number value", "./testdata/deeply_nested_flags.json", "number", "json", "example_invalid_named_number", int64(1), int64(1), "deeply", "nested"),
	Entry("json parses a float value", "./testdata/deeply_nested_flags.json", "float", "json", "example_named_float", 1.0, 2.718281828459045, "deeply", "nested"),
	Entry("json parses an invalid float value", "./testdata/deeply_nested_flags.json", "float", "json", "example_invalid_named_float", 1.0, 1.0, "deeply", "nested"),

	Entry("json parses an empty boolean value", "./testdata/empty.json", "boolean", "json", "example_invalid_boolean", false, false, "deeply", "nested"),
	Entry("json parses an empty string value", "./testdata/empty.json", "string", "json", "example_invalid_named_string", "default", "default", "deeply", "nested"),
	Entry("json parses an empty number value", "./testdata/empty.json", "number", "json", "example_invalid_named_number", int64(1), int64(1), "deeply", "nested"),
	Entry("json parses an empty float value", "./testdata/empty.json", "float", "json", "example_invalid_named_float", 1.0, 1.0, "deeply", "nested"),

	Entry("yaml parses a boolean value", "./testdata/flags.yaml", "boolean", "yaml", "example_boolean", false, true),
	Entry("yaml parses an invalid boolean value", "./testdata/flags.yaml", "boolean", "yaml", "example_invalid_boolean", false, false),
	Entry("yaml parses a string value", "./testdata/flags.yaml", "string", "yaml", "example_named_string", "default", "medium"),
	Entry("yaml parses an invalid string value", "./testdata/flags.yaml", "string", "yaml", "example_invalid_named_string", "default", "default"),
	Entry("yaml parses a number value", "./testdata/flags.yaml", "number", "yaml", "example_named_number", int64(1), int64(128)),
	Entry("yaml parses an invalid number value", "./testdata/flags.yaml", "number", "yaml", "example_invalid_named_number", int64(1), int64(1)),
	Entry("yaml parses a float value", "./testdata/flags.yaml", "float", "yaml", "example_named_float", 1.0, 2.718281828459045),
	Entry("yaml parses an invalid float value", "./testdata/flags.yaml", "float", "yaml", "example_invalid_named_float", 1.0, 1.0),

	Entry("yaml parses a boolean value", "./testdata/deeply_nested_flags.yaml", "boolean", "yaml", "example_boolean", false, true, "deeply", "nested"),
	Entry("yaml parses an invalid boolean value", "./testdata/deeply_nested_flags.yaml", "boolean", "yaml", "example_invalid_boolean", false, false, "deeply", "nested"),
	Entry("yaml parses a string value", "./testdata/deeply_nested_flags.yaml", "string", "yaml", "example_named_string", "default", "medium", "deeply", "nested"),
	Entry("yaml parses an invalid string value", "./testdata/deeply_nested_flags.yaml", "string", "yaml", "example_invalid_named_string", "default", "default", "deeply", "nested"),
	Entry("yaml parses a number value", "./testdata/deeply_nested_flags.yaml", "number", "yaml", "example_named_number", int64(1), int64(128), "deeply", "nested"),
	Entry("yaml parses an invalid number value", "./testdata/deeply_nested_flags.yaml", "number", "yaml", "example_invalid_named_number", int64(1), int64(1), "deeply", "nested"),
	Entry("yaml parses a float value", "./testdata/deeply_nested_flags.yaml", "float", "yaml", "example_named_float", 1.0, 2.718281828459045, "deeply", "nested"),
	Entry("yaml parses an invalid float value", "./testdata/deeply_nested_flags.yaml", "float", "yaml", "example_invalid_named_float", 1.0, 1.0, "deeply", "nested"),

	Entry("yaml parses an empty boolean value", "./testdata/empty.yaml", "boolean", "yaml", "example_invalid_boolean", false, false, "deeply", "nested"),
	Entry("yaml parses an empty string value", "./testdata/empty.yaml", "string", "yaml", "example_invalid_named_string", "default", "default", "deeply", "nested"),
	Entry("yaml parses an empty number value", "./testdata/empty.yaml", "number", "yaml", "example_invalid_named_number", int64(1), int64(1), "deeply", "nested"),
	Entry("yaml parses an empty float value", "./testdata/empty.yaml", "float", "yaml", "example_invalid_named_float", 1.0, 1.0, "deeply", "nested"),
)
