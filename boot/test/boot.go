package test

import (
	"github.com/spf13/viper"
)

// BootInit is a function to initialize boot
func BootInit(configFile string) {
	// load config
	// viper.SetConfigType("toml")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
