package main

import (
	"github.com/spf13/viper"
	"github.com/varz1/nCovBot/bot"
	"github.com/varz1/nCovBot/maker"
	"sync"
)

func main() {
	viper.AddConfigPath("/home/cl/go/src/github.com/varz1/nCovBot/config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("viper wrong")
	}
	go bot.Run()
	go maker.List()
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
