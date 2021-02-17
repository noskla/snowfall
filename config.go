package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	DBAddress        string `json:"dbAddress"`
	DBUser           string `json:"dbUser"`
	DBPassword       string `json:"dbPassword"`
	DBName           string `json:"dbName"`
	DBSslMode        string `json:"dbSslMode"`
	RabbitMQAddress  string `json:"rabbitMQAddress"`
	RabbitMQPort     string `json:"rabbitMQPort"`
	RabbitMQUsername string `json:"rabbitMQUsername"`
	RabbitMQPassword string `json:"rabbitMQPassword"`
	RabbitMQSecure   bool   `json:"rabbitMQSecure"`
}

func loadConfig(filename string) config {
	configFile, err := os.Open(filename)
	errorOccurred(err, true)
	defer configFile.Close()

	configFileBytes, _ := ioutil.ReadAll(configFile)
	var cfg config
	err = json.Unmarshal(configFileBytes, &cfg)
	errorOccurred(err, true)

	return cfg
}
