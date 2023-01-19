package myconfig

type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
}

type ServerConfigurations struct {
	Port int
}

type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}
