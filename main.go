package main

import (
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("/home/cl/go/src/github.com/varz1/nCovBot/config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper wrong")
	}
}
