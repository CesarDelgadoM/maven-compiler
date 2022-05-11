package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

	Info("Execution command successful")
	return nil
}
