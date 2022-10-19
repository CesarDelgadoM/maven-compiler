package main

import (
	"os"
)

func main() {

	err := InitializeApp()
	if err != nil {
		Error("Error to initialize the app")
		return
	}

	StartApp()
}

// Create the file for projects the first time the app is executed
func InitializeApp() error {

	// Get path of the current directory
	dir, err := os.Getwd()
	if err != nil {
		Error("Cannot get path directory: " + err.Error())
		os.Exit(1)
	}

	pathFile = dir + separator + fileName

	// Validates if the file exists
	_, err = os.Stat(fileName)
	if err == nil {
		readFile(pathFile)
		return nil
	}

	_, err = os.Create(pathFile)
	if err != nil {
		Error("Cannot create file [" + fileName + "]: " + err.Error())
		return err
	}

	err = hiddenFile(pathFile)
	if err != nil {
		Error("Error to hide the file")
		return err
	}

	return nil
}
