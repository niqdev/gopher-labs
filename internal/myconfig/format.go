package myconfig

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

// NOTE: lowercase fields are NOT exported!

func newMyConfig() *MyConfig {
	return &MyConfig{
		MyRoot: "foo-new",
		MyExample: MyExampleConfig{
			MyInt:     8,
			MyString:  "bar-new",
			MyDefault: "default-new",
			MyDash:    "dash-new",
		},
	}
}

func Format() {
	myConfig := newMyConfig()
	fmt.Printf("%+v\n", myConfig)
	fmt.Printf("%#v\n", myConfig)

	yamlExample(myConfig)
	jsonExample(myConfig)
}

func yamlExample(myConfig *MyConfig) {
	myConfigBytes, err := yaml.Marshal(&myConfig)
	if err != nil {
		fmt.Printf("unable to marshal config to yaml: %v", err)
	}
	myConfigString := string(myConfigBytes)
	fmt.Println(myConfigString)

	var myConfigYaml MyConfig
	err = yaml.Unmarshal([]byte(myConfigString), &myConfigYaml)
	if err != nil {
		fmt.Printf("unable to unmarshal config from yaml: %v", err)
	}
	fmt.Printf("%+v\n", myConfigYaml)
}

func jsonExample(myConfig *MyConfig) {
	myConfigBytes, err := json.MarshalIndent(&myConfig, "", "  ")
	if err != nil {
		fmt.Printf("unable to marshal config to json: %v", err)
	}
	myConfigString := string(myConfigBytes)
	fmt.Println(myConfigString)

	var myConfigJson MyConfig
	err = json.Unmarshal([]byte(myConfigString), &myConfigJson)
	if err != nil {
		fmt.Printf("unable to unmarshal config from json: %v", err)
	}
	fmt.Printf("%+v\n", myConfigJson)
}
