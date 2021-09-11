package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
    //Token export
	Token     string
    //BotPrefix export
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"Token"`
	BotPrefix string `json:"BotPrefix"`
}

//ReadConfig function
func ReadConfig() error {

	fmt.Println("\nReading config file...")

	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
