package myschema

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v3"
)

//go:embed employee.schema.json
var employeeSchema string

// EMBED FILES https://pkg.go.dev/embed
// EMBED RELATIVE PATH https://github.com/golang/go/issues/46056
// JSON TO SCHEMA https://stackoverflow.com/questions/7341537/tool-to-generate-json-schema-from-json-data

// EXAMPLE https://www.mongodb.com/basics/json-schema-examples
// https://github.com/niqdev/kotlin-fun/blob/main/modules/json-schema/src/test/kotlin/com/github/niqdev/json/ExampleTest.kt
// cat data/employee.json | yq -p json -o yaml > data/employee.yaml
func JsonSchemaValidation() {

	// removes timestamps
	log.SetFlags(0)

	fmt.Println(employeeSchema)

	fmt.Println("LIB: xeipuuv/gojsonschema")
	goJsonSchemaExample("employee.json")
	goJsonSchemaExample("employee-invalid.json")

	fmt.Println("LIB: santhosh-tekuri/jsonschema")
	jsonSchemaExampleJson("employee.json")
	jsonSchemaExampleJson("employee-invalid.json")
	jsonSchemaExampleYaml("employee.yaml")
	jsonSchemaExampleYaml("employee-invalid.yaml")
}

// https://github.com/xeipuuv/gojsonschema
// https://github.com/docker/libcompose/blob/master/config/schema.go
// https://github.com/docker/libcompose/blob/master/config/validation.go
// YAML wrapper https://github.com/neilpa/yajsv
func goJsonSchemaExample(fileName string) {

	schemaLoader := gojsonschema.NewStringLoader(employeeSchema)
	validDocumentLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://data/%s", fileName))

	result, err := gojsonschema.Validate(schemaLoader, validDocumentLoader)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		fmt.Printf("The document [%s] is valid\n", fileName)
	} else {
		fmt.Printf("The document [%s] is not valid. See errors:\n", fileName)
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}

// https://github.com/santhosh-tekuri/jsonschema/blob/master/example_test.go
func jsonSchemaExampleJson(fileName string) {

	schema, err := jsonschema.CompileString("schema.json", employeeSchema)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	data, err := ioutil.ReadFile(fmt.Sprintf("data/%s", fileName))
	if err != nil {
		log.Fatal(err)
	}

	// TODO EmployeeJson
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		log.Fatal(err)
	}

	// validate
	if err = schema.Validate(v); err != nil {
		log.Println(fmt.Sprintf("%#v", err))
	} else {
		log.Println("valid")
	}
}

// https://github.com/santhosh-tekuri/jsonschema/issues/5
// https://github.com/santhosh-tekuri/jsonschema/blob/master/cmd/jv/main.go
func jsonSchemaExampleYaml(fileName string) {

	// VERSION >>> santhosh-tekuri/jsonschema/v5

	// loads schema
	schema, err := jsonschema.CompileString("schema.json", employeeSchema)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	// loads data
	filePath, _ := filepath.Abs(fmt.Sprintf("data/%s", fileName))
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("%#v", err)
	}
	// TODO EmployeeYaml
	var model interface{}
	if err := yaml.Unmarshal([]byte(data), &model); err != nil {
		log.Fatalf("%#v", err)
	}

	// detailed error
	if ve, ok := schema.Validate(model).(*jsonschema.ValidationError); ok {
		// ve.FlagOutput() or ve.BasicOutput()
		b, _ := json.MarshalIndent(ve.DetailedOutput(), "", "  ")
		fmt.Fprintln(os.Stderr, string(b))
	} else {
		log.Println("valid")
	}
}
