package myschema

import (
	_ "embed"
	"fmt"
)

//go:embed employee.schema.json
var employeeSchema string

// EMBED FILES https://pkg.go.dev/embed
// EMBED RELATIVE PATH https://github.com/golang/go/issues/46056
// JSON TO SCHEMA https://stackoverflow.com/questions/7341537/tool-to-generate-json-schema-from-json-data

// EXAMPLE https://www.mongodb.com/basics/json-schema-examples
// https://github.com/niqdev/kotlin-fun/blob/main/modules/json-schema/src/test/kotlin/com/github/niqdev/json/ExampleTest.kt

// TODO
func JsonSchemaValidation() {
	fmt.Print(employeeSchema)
}
