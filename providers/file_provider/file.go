package file_provider

import (
	"encoding/json"

	"fmt"
	"os"

	"github.com/ehowe/openfeature-go-providers/shared"
	"github.com/open-feature/go-sdk/openfeature"
	"gopkg.in/yaml.v3"
)

type file struct {
	contents []shared.Flag
	path     string
	format   string
	deepKeys []string
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (f *file) readFile() error {
	if len(f.contents) > 0 {
		return nil
	}

	if !contains([]string{"json", "yaml"}, f.format) {
		return fmt.Errorf("Unknown format: %s", f.format)
	}

	file, err := os.Open(f.path)

	if err != nil {
		return err
	}

	defer file.Close()

	var rawFlags interface{}
	if f.format == "json" {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&rawFlags)
	} else {
		decoder := yaml.NewDecoder(file)
		err = decoder.Decode(&rawFlags)
	}

	if err != nil {
		return err
	}

	nestedContent := make(map[string]interface{})
	var nestedContents []interface{}

	if len(f.deepKeys) > 0 {
		for i, key := range f.deepKeys {
			if i == 0 {
				nestedContent = rawFlags.(map[string]interface{})[key].(map[string]interface{})
			} else if i == len(f.deepKeys)-1 {
				nestedContents = nestedContent[key].([]interface{})
			} else {
				nestedContent = nestedContent[key].(map[string]interface{})
			}
		}
	} else {
		nestedContents = rawFlags.([]interface{})
	}

	var contents []shared.Flag

	for _, content := range nestedContents {
		var contentStruct shared.Flag
		jsonData, _ := json.Marshal(content)
		err = json.Unmarshal(jsonData, &contentStruct)

		if err != nil {
			return err
		}

		contents = append(contents, contentStruct)
	}

	f.contents = contents

	return err
}

func (f *file) fetchFlag(key string) (shared.Flag, error) {
	var resolvedFlag shared.Flag
	err := f.readFile()

	if err != nil {
		panic(fmt.Sprintf("Unable to parse file %s: %s", f.path, err))
	}

	for _, flag := range f.contents {
		if flag.Name == key {
			resolvedFlag = flag
		}
	}

	if resolvedFlag.Name != "" {
		return resolvedFlag, nil
	}
	return resolvedFlag, openfeature.NewFlagNotFoundResolutionError("")
}
