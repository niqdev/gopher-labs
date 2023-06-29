package myconfig

type MyConfig struct {
	MyRoot    string
	MyExample MyExampleConfig
}

type MyExampleConfig struct {
	MyInt     int
	MyString  string
	MyDefault string
	MyDash    string `mapstructure:"my-dash"` // dash issue https://stackoverflow.com/questions/51228423/how-to-unmarshall-viper-config-to-struct-with-dash-character
}
