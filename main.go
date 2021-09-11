package main

import (
	//"fmt"
	"./bot"
	"./config"
)

func main() {

	err := config.ReadConfig()
	if err != nil {
		println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
