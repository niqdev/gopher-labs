package myschema

import (
	_ "embed"
	"fmt"
)

//go:embed example.json
var exampleSchema string

// EMBED FILES https://pkg.go.dev/embed
// EMBED RELATIVE PATH https://github.com/golang/go/issues/46056
// JSON TO SCHEMA https://stackoverflow.com/questions/7341537/tool-to-generate-json-schema-from-json-data

func JsonSchemaValidation() {
	fmt.Print(exampleSchema)
}
