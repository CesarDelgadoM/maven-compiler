package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CompileProject() {

	keys := projects.Keys()

	op := InputOption("compile", keys)
	if strings.EqualFold(op, "0") {
		Info("Saliendo...")
		return
	}

	paths := projects.Value(op)

	var profile string
	if op := InputText("Compilation profile?(Y/n): "); strings.ToLower(op) == "y" {
		fmt.Print("Profile name: ")
		fmt.Scanln(&profile)
	}

	size := len(paths)
	for i := 0; i < size; i++ {

		cls()
		Info("Compiling the project: " + paths[i])

		var err error
		// Check if it is the last project and if compiles with profile
		if i == size-1 && profile != "" {
			err = Compile(paths[i], profile)
		} else {
			err = Compile(paths[i], "")
		}

		if err != nil {
			Error("project compilation failed: " + err.Error())
			return
		}
	}
	Info("Compilation of the project successfully")
	PressEnter()
}

// Compiles projects of maven type
func Compile(pathProject string, profile string) error {
	err := os.Chdir(pathProject)
	if err != nil {
		Error(err.Error())
	}

	var command string

	if profile == "" {
		command = "/c mvn clean compile package install"
	} else {
		command = "/c mvn clean compile package install -P " + profile
	}
	err = ExecuteCommand(command)
	return err
}

// Execute commands of the windows system
func ExecuteCommand(command string) error {
	cmd := exec.Command("cmd", strings.Split(command, " ")...)

	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	if err = cmd.Wait(); err != nil {
		return err
	}

	Info("Execution command successful")
	return nil
}
