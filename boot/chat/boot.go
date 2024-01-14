package chat

import (
	"flag"

	"github.com/rumis/ray"
	"github.com/spf13/viper"
)

var configFile *string

// BootInit is a function to initialize boot
func BootInit() {

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

	ray.SetDefaultProxy(viper.GetString("proxy.addr")) // set proxy
	ray.SetDefaultRetryTimesAndTimeout(60, 2)          // set retry times and timeout

}
