package main

import (
	"log"
	"math/rand"
	"os"
)

func errorOccurred(err error, fatal bool) bool {
	if err != nil {
		log.Println(err.Error())
		if fatal {
			log.Fatalln("Halting program.")
		}
		return true
	}
	return false
}

func getDirectoryPath(name string) string {
	path := Here + "/" + name
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// try finding templates relatively
		path = name
		if _, err = os.Stat(path); os.IsNotExist(err) {
			log.Fatalln("Could not find", name, "directory.")
		}
	}
	return path
}

var runePool = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
func getRandomString(length int) string {
	randStr := make([]rune, length)
	for i := range randStr {
		randStr[i] = runePool[rand.Intn(len(runePool))]
	}
	return string(randStr)
}
