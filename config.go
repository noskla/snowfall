package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	DBAddress  string `json:"dbAddress"`
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
	DBSslMode  string `json:"dbSslMode"`
}

func loadConfig(filename string) config {
	configPath := getDirectoryPath(filename)
	configFile, err := os.Open(configPath)
	errorOccurred(err, true)
	defer configFile.Close()

	configFileBytes, _ := ioutil.ReadAll(configFile)
	var cfg config
	err = json.Unmarshal(configFileBytes, &cfg)
	errorOccurred(err, true)

	return cfg
}