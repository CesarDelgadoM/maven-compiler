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
	name, exit := InputOption("compile", keys)
	if exit {
		return
	}

	paths := projects.Value(name)

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
			err = ExecuteCommandMaven(paths[i], profile)
		} else {
			err = ExecuteCommandMaven(paths[i], "")
		}

		if err != nil {
			Error("Compilation of the project", name, "failed:", err)
			PressEnter()
			return
		}
	}
	Info("Compilation of the project", name, "successfully")
	PressEnter()
}

// Execute commands of projects maven
func ExecuteCommandMaven(pathProject string, profile string) error {
	err := os.Chdir(pathProject)
	if err != nil {
		Error(err.Error())
		return err
	}

	var command string

	if profile == "" {
		command = "/c mvn clean compile package install"
	} else {
		command = "/c mvn clean compile package install -P " + profile
	}
	return ExecuteCommand(command)
}

// Execute commands of the windows system
func ExecuteCommand(command string) error {

	cmd := exec.Command("cmd", strings.Split(command, " ")...)

	out, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}
