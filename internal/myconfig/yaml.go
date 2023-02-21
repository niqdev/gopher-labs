package myconfig

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func newMyConfig() *MyConfig {
	return &MyConfig{
		MyRoot: "foo-new",
		MyExample: MyExampleConfig{
			MyInt:     8,
			MyString:  "bar-new",
			MyDefault: "default-new",
		},
	}
}

func YamlExample() {
	myConfig := newMyConfig()
	fmt.Printf("%+v\n", myConfig)
	fmt.Printf("%#v\n", myConfig)

	myConfigBytes, err := yaml.Marshal(&myConfig)
	if err != nil {
		fmt.Printf("unable to marshal config to yaml: %v", err)
	}
	myConfigString := string(myConfigBytes)
	fmt.Println(myConfigString)

	var myConfigNew MyConfig
	err = yaml.Unmarshal([]byte(myConfigString), &myConfigNew)
	if err != nil {
		fmt.Printf("unable to unmarshal config from yaml: %v", err)
	}
	fmt.Printf("%+v\n", myConfigNew)
}
