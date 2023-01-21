package myconfig

type MyConfig struct {
	MyRoot    string
	MyExample MyExampleConfig
}

type MyExampleConfig struct {
	MyInt     int
	MyString  string
	MyDefault string
}
