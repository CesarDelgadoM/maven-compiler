package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	separator string = "\\"
	fileName  string = ".projects-paths.json"
)

var pathFile string
var projects project = make(project, 100)

func MenuFiles() {

	var op int
	var exit bool
	for !exit {
		MenuFilesOptions()
		fmt.Print("[option] :: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			CrateProject()
		case 2:
			OpenProject()
		case 3:
			EditProject()
		case 4:
			DeleteProject()
		case 5:
			ListNamesProjects()
		case 0:
			Info("Saliendo...")
			exit = true
		default:
			Info("Option not exist, try again!")
		}
	}
}

func CrateProject() {

	name := InputText("Type the name of the project: ")
	PrintLine()

	paths := InputPaths()
	fmt.Println(paths)
	projects[name] = paths

	defer WriteFile(pathFile, projects)
}

func OpenProject() {

	keys := projects.Keys()

	op := InputOption("open", keys)
	if strings.EqualFold(op, "0") {
		Info("Saliendo...")
		return
	}

	PrintLine()
	paths := projects.Value(op)
	for p := range paths {
		fmt.Println("->", p)
	}
	PrintLine()
	PressEnter()
}

func EditProject() {

	keys := projects.Keys()

	op := InputOption("edit", keys)
	if strings.EqualFold(op, "0") {
		Info("Saliendo...")
		return
	}

	paths := InputPaths()
	projects[op] = paths

	defer WriteFile(pathFile, projects)
}

func DeleteProject() {

	keys := projects.Keys()

	op := InputOption("name", keys)
	if strings.EqualFold(op, "0") {
		Info("Saliendo...")
		return
	}

	delete(projects, op)

	defer WriteFile(pathFile, projects)
}

func InputPaths() []string {

	fmt.Println("En caso de ser un solo proyecto ingrese la ruta, de lo contrario ingrese las rutas en orden de compilacion separadas por (,): ")
	fmt.Print("Paths: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pathsProjects := scanner.Text()

	return strings.Split(pathsProjects, ",")
}
