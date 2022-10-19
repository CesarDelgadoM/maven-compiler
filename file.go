package main

import (
	"encoding/json"
	"os"
	"syscall"
)

func readFile(path string) {

	file, err := os.ReadFile(path)
	if err != nil {
		Error("Error to read file")
		return
	}

	err = json.Unmarshal([]byte(file), &projects)
	if err != nil {
		Error(err.Error())
		return
	}
}

func hiddenFile(path string) error {

	fileNameW, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		Error("Error to convert path to UTF16")
		return err
	}

	//Hides the file
	err = syscall.SetFileAttributes(fileNameW, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		Error("Error to set file attribute hidden")
		return err
	}

	return nil
}

func WriteFile(path string, data []project) {

	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		Error(err.Error())
		return
	}

	err = os.WriteFile(path, file, os.ModePerm)
	if err != nil {
		Error(err.Error())
		return
	}
}
