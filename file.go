package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	nameFolder string = "compiler_files"
	formatFile string = ".txt"
	separator  string = "\\"
)

var currentPath string
var pathFolder string
var files []string

func MenuFiles() {

	var op int = -1
	for op != 0 {
		MenuFilesOptions()
		fmt.Print("[opcion] :: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			CreateCompilerFile()
		case 2:
			OpenCompilerFile()
		case 3:
			EditCompilerFile()
		case 4:
			DeleteCompilerFile()
		case 5:
			ListCompilerFiles()
		case 0:
			Info("Saliendo...")
			return
		default:
			Info("Option not exist, try again!")
		}
		PressEnter()
	}
}

func CreateCompilerFile() {

	fileName := InputText("Digite el nombre del proyecto: ")
	PrintLine()

	path := pathFolder + separator + fileName + formatFile

	file, err := os.Create(path)
	if err != nil {
		Error("Cannot create file [" + fileName + "]: " + err.Error())
		return
	}

	var pathsProjects string
	fmt.Println("En caso de ser un solo proyecto ingrese la ruta, de lo contrario ingrese las rutas en orden de compilacion separadas por (,): ")
	fmt.Print("Rutas: ")
	fmt.Scanln(&pathsProjects)

	file.WriteString(pathsProjects)
	files = append(files, fileName+formatFile)
	file.Close()

	Info("Archivo creado con exito!")
}

func EditCompilerFile() {

	CreateMenu(files)
	op := InputNumber("[editar] :: ")
	if op <= 0 || op > len(files) {
		if op == 0 {
			Info("Saliendo...")
		} else {
			Error("Option index out of bounds")
		}
		return
	}
	fileName := files[op-1]

	path := pathFolder + separator + fileName

	file, err := os.OpenFile(path, os.O_RDWR, 0600)
	if err != nil {
		Error("Cannot open file [" + fileName + "]: " + err.Error())
		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		Error("Cannot read the file: " + err.Error())
		return
	}
	Info("Rutas del proyecto: [" + string(content) + "]")

	var pathsProjects string
	fmt.Println("En caso de ser un solo proyecto ingrese la ruta, de lo contrario ingrese las rutas en orden de compilacion separadas por (,): ")
	fmt.Print("Rutas: ")
	fmt.Scanln(&pathsProjects)

	if pathsProjects != "" {
		err = file.Truncate(0)
		_, err = file.WriteString(pathsProjects)
		if err != nil {
			Error("Cannot edit file: " + err.Error())
		}
		Info("File edited successful")
	} else {
		Info("File not edited, no data")
	}
	file.Close()
}

func OpenCompilerFile() {

	CreateMenu(files)

	op := InputNumber("[abrir] :: ")
	if op <= 0 || op > len(files) {
		if op == 0 {
			Info("Saliendo...")
		} else {
			Error("Option index out of bounds")
		}
		return
	}

	fileName := files[op-1]
	path := pathFolder + separator + fileName

	content, err := os.ReadFile(path)
	if err != nil {
		Error("Cannot read the file: " + err.Error())
		return
	}

	PrintLine()
	fmt.Println(string(content))
	PrintLine()
}

func DeleteCompilerFile() {

	CreateMenu(files)

	op := InputNumber("[eliminar] :: ")
	if op <= 0 || op > len(files) {
		if op == 0 {
			Info("Saliendo...")
		} else {
			Error("Option index out of bounds")
		}
		return
	}
	fileName := files[op-1]

	path := pathFolder + separator + fileName

	err := os.Remove(path)
	if err != nil {
		Error("Cannot delete file: " + err.Error())
	} else {
		ReadDirectory()
		Info("File deleted successful")
	}
}

func ListCompilerFiles() {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                         [ARCHIVOS COMPILACION]                         ")
	fmt.Println("------------------------------------------------------------------------")
	for _, file := range files {
		fmt.Println("->[" + file + "]")
	}
	PrintLine()
}

func InitializeApp() {
	// Get path of the current directory on the variable currentPath
	path, err := os.Getwd()
	if err != nil {
		Error("Cannot get path directory: " + err.Error())
		os.Exit(1)
	}
	currentPath = path

	// Create folder if not exist
	pathFolder = currentPath + separator + nameFolder
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			Error("Cannot create folder: " + err.Error())
			os.Exit(1)
		}
	}

	ReadDirectory()
}

// Reads the directory to fill the list of compiler files
func ReadDirectory() {
	filesInfo, err := ioutil.ReadDir(pathFolder)
	if err != nil {
		Error("Cannot read the directory: " + err.Error())
		return
	}

	files = nil
	for _, file := range filesInfo {
		if strings.Contains(file.Name(), ".txt") {
			files = append(files, file.Name())
		}
	}
}

func UpdateCurrentPath() {
	err := os.Chdir(currentPath)
	if err != nil {
		Error("Cannot update directory path: " + err.Error())
		os.Exit(1)
	}
}
