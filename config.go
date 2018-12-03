package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ConfigFileName は config.json のファイル名
var ConfigFileName = "config.json"

// Config には各キーが格納される
type Config struct {
	APIKey            string `json:"API key"`
	APISecretKey      string `json:"API secret key"`
	AccessToken       string `json:"Access token"`
	AccessTokenSecret string `json:"Access token secret"`
}

// ReadConfigFile は config.json ファイルを読み込む
func ReadConfigFile() *Config {
	bytes, err := ioutil.ReadFile(ConfigFileName)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatal(err)
	}

	return &config
}
