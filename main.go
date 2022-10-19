package main

import (
	"fmt"
	"os"
)

func main() {

	err := InitializeApp()
	if err != nil {
		Error("Error to initialize the app")
		return
	}

	var op int
	for {
		Menu()
		fmt.Print("[option] :: ")
		fmt.Scanln(&op)
		PrintLine()

		switch op {
		case 1:
			CompileProject()
		case 2:
			MenuFiles()
			continue
		case 0:
			Info("Saliendo del compilador...")
			os.Exit(0) // No-Error
		default:
			Info("Option not exist, try again!")
		}
	}
}

func InitializeApp() error {

	// Get path of the current directory on the variable currentPath
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
