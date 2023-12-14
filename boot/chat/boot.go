package chat

import (
	"flag"

	"github.com/rumis/ray"
	"github.com/spf13/viper"
)

var configFile *string

// BootInit is a function to initialize boot
func BootInit() {

	ray.SetDefaultRetryTimesAndTimeout(10, 2)

	// parse args
	configFile = flag.String("config", "config.toml", "config file path")
	flag.Parse()

	// load config
	// viper.SetConfigType("toml")
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
