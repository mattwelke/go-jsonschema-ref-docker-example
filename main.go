package main

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

func main() {
	schemaLoader := gojsonschema.NewReferenceLoader("file://./schemas/person.schema.json")

	jsonStr := `
	{
		"name": "Matt",
		"pet": {
			"name": "Shady"
		}
	}
	`

	documentLoader := gojsonschema.NewStringLoader(jsonStr)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(fmt.Errorf("could not validate: %w", err))
	}

	if result.Valid() {
		fmt.Printf("The document is valid.\n")
	} else {
		fmt.Printf("The document is not valid. See errors:\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
