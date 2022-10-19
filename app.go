package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	separator string = "\\"
	fileName  string = ".projects-routes.json"
)

var pathFile string

type project map[string][]string

var projects []project

func MenuFiles() {

	var op int = -1
	for op != 0 {
		MenuFilesOptions()
		fmt.Print("[opcion] :: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			CrateCompilationProject()
		case 0:
			Info("Saliendo...")
			return
		default:
			Info("Option not exist, try again!")
		}
		PressEnter()
		op = -1
	}
}

func CrateCompilationProject() {

	name := InputText("Type the name of the project: ")
	PrintLine()

	paths := InputPaths()

	projects = append(projects, project{name: paths})

	WriteFile(pathFile, projects)
}

func InputPaths() (paths []string) {

	fmt.Println("En caso de ser un solo proyecto ingrese la ruta, de lo contrario ingrese las rutas en orden de compilacion separadas por (,): ")
	fmt.Print("Rutas: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pathsProjects := scanner.Text()

	paths = strings.Split(pathsProjects, ",")

	return
}

// func ListCompilerFiles() {
// 	cls()
// 	fmt.Println("------------------------------------------------------------------------")
// 	fmt.Println("                         [ARCHIVOS COMPILACION]                         ")
// 	fmt.Println("------------------------------------------------------------------------")
// 	for _, file := range files {
// 		fmt.Println("->[" + file + "]")
// 	}
// 	PrintLine()
// }

func InitializeApp() error {

	// Get path of the current directory on the variable currentPath
	path, err := os.Getwd()
	if err != nil {
		Error("Cannot get path directory: " + err.Error())
		os.Exit(1)
	}

	pathFile = path + separator + fileName

	_, err = os.Stat(fileName)
	if err == nil {
		Info("File Exist!")
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
