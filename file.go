package main

import (
	"encoding/json"
	"os"
	"syscall"
)

func readFile(path string) {

	file, err := os.ReadFile(path)
	if err != nil {
		Error("Error to read the file")
		return
	}

	err = json.Unmarshal([]byte(file), &projects)
	if err != nil {
		Error(err)
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

func WriteFile(path string, data project) {

	datajson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		Error(err)
		return
	}

	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		Error(err)
		return
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		Error(err)
		return
	}

	_, err = file.Write(datajson)
	if err != nil {
		Error(err)
		return
	}
}
