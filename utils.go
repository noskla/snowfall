package main

import (
	"log"
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
