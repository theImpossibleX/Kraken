package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Threads       int    `json:"threads"`
	Prefix        string `json:"prefix"`
	CheckURL      string `json:"checkURL"`
	SuccessKey    string `json:"successKey"`
	ProxyFilepath string `json:"proxy-filepath"`
	Timeout       int    `json:"timeout"`
	Debug         bool   `json:"debug"`
	OutputFolder  string ""
}

var GlobalConfig Config

func Load(filepath string) {
	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {

		log.Fatalf("Error reading config file: %v", err)
	}
	err = json.Unmarshal(configFile, &GlobalConfig)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}
