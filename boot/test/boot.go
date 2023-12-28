package test

import (
	"runtime"

	"github.com/rumis/ray"
	"github.com/spf13/viper"
)

// BootInit is a function to initialize boot
func BootInit(configFile string) {

	runtime.GOMAXPROCS(8) // set max cpu

	// init request lib
	ray.SetDefaultProxy("http://127.0.0.1:10808") // set proxy
	ray.SetDefaultRetryTimesAndTimeout(60, 2)     // set retry times and timeout

	// load config
	// viper.SetConfigType("toml")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
