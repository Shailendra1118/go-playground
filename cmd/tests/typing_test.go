package tests

import (
	"testing"
	"fmt"
	"os"
	//"log/slog"

	"gopkg.in/yaml.v3"
)

func TestXxx(t *testing.T) {
	//var typing bool
	typing := false
	fmt.Printf("boolean %t\n", typing)

	fmt.Printf("character: %c\n", 33)
	fmt.Printf("integer: %d\n", 33)
	fmt.Printf("integer: %T\n", 33)

	fmt.Printf("character: %c\n", 512)
	
	//print string
	fmt.Printf("strings: %s\n", "Shailendra")
	fmt.Printf("strings: %q\n", "Shailendra") //with quotes	
}

func TestStructInterfaceMap(t *testing.T) {
	data, err := os.ReadFile("/Users/ssyadav/Documents/go-workspace/src/go-playground/cmd/tests/schema.yaml")
    if err != nil {
        fmt.Printf("Failed to read file: %v\n", err)
    }

	var yamlContent map[string]interface{}
	if err := yaml.Unmarshal(data, &yamlContent); err != nil {
		fmt.Printf("Failed to umarshall YAML file: %v\n", err)
	}

	fmt.Printf("OKay to umarshall YAML file: %v\n", yamlContent)
	version := yamlContent["apiVersion"].(string)
	
	fmt.Printf("apiVersion: %v", version)

}

