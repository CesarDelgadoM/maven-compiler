package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func Menu() {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                           [COMPILADOR MAVEN]                           ")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[1] :: [Compilar un proyecto]")
	fmt.Println("[2] :: [Menu archivos de compilacion]")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[0] :: [Salir]")
	fmt.Println("------------------------------------------------------------------------")
}

func MenuFilesOptions() {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                           [MENU  ARCHIVOS]                             ")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[1] :: [Crear archivo de compilacion]")
	fmt.Println("[2] :: [Abrir archivo de compilacion]")
	fmt.Println("[3] :: [Editar archivo de compilacion]")
	fmt.Println("[4] :: [Eliminar archivo de compilacion]")
	fmt.Println("[5] :: [Ver archivos de compilacion]")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[0] :: [Volver]")
	fmt.Println("------------------------------------------------------------------------")
}

func CreateMenu(items []string) {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                          [ARCHIVOS PROYECTOS]                          ")
	fmt.Println("------------------------------------------------------------------------")
	for i, item := range items {
		fmt.Println("[" + strconv.Itoa(i+1) + "] :: [" + item + "]")
	}
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[0] :: [Volver]")
	fmt.Println("------------------------------------------------------------------------")
}

func InputNumber(text string) int {
	var op int
	fmt.Print(text)
	fmt.Scanln(&op)
	return op
}

func InputText(text string) string {
	var op string
	fmt.Print(text)
	fmt.Scanln(&op)
	return op
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PressEnter() {
	var enter int
	fmt.Print("prees enter to continue...")
	fmt.Scanln(&enter)
}

func Info(msg string) {
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[INFO]", msg)
	fmt.Println("------------------------------------------------------------------------")
}

func Error(msg string) {
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[ERROR]", msg)
	fmt.Println("------------------------------------------------------------------------")
}

func PrintLine() {
	fmt.Println("------------------------------------------------------------------------")
}
