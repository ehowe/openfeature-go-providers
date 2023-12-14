package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var flagsString = `
[
  {
    "name": "example_boolean",
    "kind": "boolean",
    "value": true,
    "enabled": true
  },
  {
    "name": "example_invalid_boolean",
    "kind": "boolean",
    "value": "a string",
    "enabled": true
  },
  {
    "name": "example_disabled_boolean",
    "kind": "boolean",
    "value": true,
    "enabled": false
  },
  {
    "name": "example_named_string",
    "kind": "string",
    "value": "medium",
    "enabled": true
  },
  {
    "name": "example_invalid_named_string",
    "kind": "string",
    "value": true,
    "enabled": true
  },
  {
    "name": "example_disabled_named_string",
    "kind": "string",
    "variant": "medium",
    "enabled": false
  },
  {
    "name": "example_named_number",
    "kind": "number",
    "variants": {
      "small": 8,
      "medium": 128,
      "large": 2048
    },
    "variant": "medium",
    "enabled": true
  },
  {
    "name": "example_invalid_named_number",
    "kind": "number",
    "variants": {
      "small": "a string",
      "medium": "a string",
      "large": "a string"
    },
    "variant": "medium",
    "enabled": true
  },
  {
    "name": "example_disabled_named_number",
    "kind": "number",
    "variants": {
      "small": 8,
      "medium": 128,
      "large": 2048
    },
    "variant": "medium",
    "enabled": false
  },
  {
    "name": "example_named_float",
    "kind": "float",
    "variants": {
      "pi": 3.141592653589793,
      "e": 2.718281828459045,
      "phi": 1.618033988749894
    },
    "variant": "e",
    "enabled": true
  },
  {
    "name": "example_invalid_named_float",
    "kind": "float",
    "variants": {
      "pi": "a string",
      "e": "a string",
      "phi": "a string"
    },
    "variant": "e",
    "enabled": true
  },
  {
    "name": "example_disabled_named_float",
    "kind": "float",
    "variants": {
      "pi": 3.141592653589793,
      "e": 2.718281828459045,
      "phi": 1.618033988749894
    },
    "variant": "e",
    "enabled": false
  }
]
`

func writeTestData(data []byte, filename string) {
	fullPath := fmt.Sprintf("./providers/file_provider/testdata/%s", filename)

	err := os.WriteFile(fullPath, []byte(data), 0644)

	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println(flagsString)
	var rawFlags []map[string]interface{}
	err := json.Unmarshal([]byte(flagsString), &rawFlags)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", rawFlags)
	deepFlags := map[string]map[string]interface{}{"deeply": {"nested": rawFlags}}
	prettyJson, _ := json.MarshalIndent(rawFlags, "", "  ")
	prettyJsonDeep, _ := json.MarshalIndent(deepFlags, "", "  ")
	yamlFlags, _ := yaml.Marshal(rawFlags)
	deepYaml, _ := yaml.Marshal(deepFlags)
	writeTestData(prettyJson, "flags.json")
	writeTestData(prettyJsonDeep, "deeply_nested_flags.json")
	writeTestData(yamlFlags, "flags.yaml")
	writeTestData(deepYaml, "deeply_nested_flags.yaml")
}
