package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	InitializeApp()

	var op int = -1
	for op != 0 {
		Menu()
		fmt.Print("[opcion] :: ")
		fmt.Scanln(&op)
		PrintLine()

		switch op {
		case 1:
			CompileProject()
			UpdateCurrentPath()
		case 2:
			MenuFiles()
			continue
		case 0:
			Info("Saliendo del compilador...")
			return
		default:
			Info("Option not exist, try again!")
		}
		PressEnter()
		op = -1
	}
}

func CompileProject() {

	op1 := InputOption("compilar", files)
	if op1 == 0 {
		Info("Saliendo...")
		return
	}

	fileName := files[op1-1]
	path := pathFolder + separator + fileName

	content, err := os.ReadFile(path)
	if err != nil {
		Error("Cannot read the file: " + err.Error())
		return
	}

	op2 := InputText("Perfil de compilacion?(S/n): ")

	var profile string
	if strings.ToLower(op2) == "s" {
		fmt.Print("Nombre perfil: ")
		fmt.Scanln(&profile)
	}

	pathsProjects := strings.Split(string(content), ",")
	size := len(pathsProjects)

	for i := 0; i < size; i++ {
		cls()
		Info("Compilando el proyecto: " + pathsProjects[i])

		if i == size-1 && profile != "" {
			err = Compile(pathsProjects[i], profile)
		} else {
			err = Compile(pathsProjects[i], "")
		}
		if err != nil {
			Error("project compilation failed: " + err.Error())
			return
		}
	}
	Info("Compilation of the project successfully")
}
