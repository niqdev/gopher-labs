package myschema

// https://stackoverflow.com/questions/28682439/go-parse-yaml-file

// https://mholt.github.io/json-to-go
type EmployeeJson struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Hobbies struct {
		Indoor  []string `json:"indoor"`
		Outdoor []string `json:"outdoor"`
	} `json:"hobbies"`
}

// https://zhwt.github.io/yaml-to-go
type EmployeeYaml struct {
	ID      int    `yaml:"id"`
	Name    string `yaml:"name"`
	Age     int    `yaml:"age"`
	Hobbies struct {
		Indoor  []string `yaml:"indoor"`
		Outdoor []string `yaml:"outdoor"`
	} `yaml:"hobbies"`
}
