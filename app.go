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

func StartApp() {
	for {
		MenuCompiler()
		op := InputNumber("[option] :: ")

		switch op {
		case 1:
			CompileProject()
		case 2:
			MenuProjects()
		case 0:
			Exit()
		default:
			OptionNotExist()
			continue
		}
	}
}

func MenuProjects() {

	var op int
	for {
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
			return
		default:
			OptionNotExist()
			continue
		}
	}
}

func CrateProject() {

	name := InputText("Type the name of the project: ")
	PrintLine()

	paths := InputPaths()
	projects[name] = paths

	Info("Project create successfully")
	PressEnter()

	defer WriteFile(pathFile, projects)
}

func OpenProject() {

	keys := projects.Keys()
	op, back := InputOption("open", keys)
	if back {
		return
	}

	PrintLine()
	paths := projects.Value(op)
	for _, p := range paths {
		fmt.Println("->", p)
	}
	PrintLine()
	PressEnter()
}

func EditProject() {

	keys := projects.Keys()
	op, back := InputOption("edit", keys)
	if back {
		return
	}

	paths := InputPaths()
	projects[op] = paths

	Info("Project updated successfully")
	PressEnter()

	defer WriteFile(pathFile, projects)
}

func DeleteProject() {

	keys := projects.Keys()
	op, back := InputOption("delete", keys)
	if back {
		return
	}

	delete(projects, op)

	Info("Project deleted successfully")
	PressEnter()

	defer WriteFile(pathFile, projects)
}

func InputPaths() []string {

	fmt.Println("Enter project paths in order and separated by (,): ")
	fmt.Print("Paths: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	pathsProjects := scanner.Text()

	return strings.Split(pathsProjects, ",")
}
