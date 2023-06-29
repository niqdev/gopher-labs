package myconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func Load() {
	// name of config file (without extension)
	viper.SetConfigName("myconfig")

	// REQUIRED if the config file does not have the extension in the name
	viper.SetConfigType("yaml")

	// path to look for the config file in
	// call multiple times to add many search paths
	viper.AddConfigPath("./data")
	viper.AddConfigPath(".")

	// env variables are case sensitive
	viper.AutomaticEnv()
	viper.SetEnvPrefix("tmp")
	// keys are automatically uppercased and prefixed
	os.Setenv("TMP_MYEXAMPLE.MYINT", "8")

	viper.SetDefault("myexample.mydefault", "DEFAULT")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("error reading config: %w", err))
	}

	log.Println(fmt.Sprintf("myroot=%s", viper.GetString("myroot")))
	log.Println(fmt.Sprintf("myexample.myint=%d", viper.GetInt("myexample.myint")))
	log.Println(fmt.Sprintf("myexample.mystring=%s", viper.GetString("myexample.mystring")))
	log.Println(fmt.Sprintf("myexample.mydefault=%s", viper.GetString("myexample.mydefault")))
	log.Println(fmt.Sprintf("myexample.mydefault=%s", viper.GetString("myexample.mydefault")))
	log.Println(fmt.Sprintf("myexample.my-dash=%s", viper.GetString("myexample.my-dash")))

	var myConfig MyConfig
	if err := viper.Unmarshal(&myConfig); err != nil {
		log.Fatal(fmt.Errorf("error decoding config: %w", err))
	}

	log.Println(fmt.Sprintf("myroot=%s", myConfig.MyRoot))
	log.Println(fmt.Sprintf("myexample.myint=%d", myConfig.MyExample.MyInt))
	log.Println(fmt.Sprintf("myexample.mystring=%s", myConfig.MyExample.MyString))
	log.Println(fmt.Sprintf("myexample.mydefault=%s", myConfig.MyExample.MyDefault))
	log.Println(fmt.Sprintf("myexample.my-dash=%s", myConfig.MyExample.MyDash))
}
