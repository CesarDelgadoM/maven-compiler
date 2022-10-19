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
	fmt.Println("                           [MAVEN COMPILER]                           ")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[1] :: [Compile a project]")
	fmt.Println("[2] :: [Menu]")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[0] :: [Exit]")
	fmt.Println("------------------------------------------------------------------------")
}

func MenuFilesOptions() {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                           [MENU  ARCHIVOS]                             ")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[2] :: [Open project routes]")
	fmt.Println("[3] :: [Edit project routes]")
	fmt.Println("[4] :: [Delete project]")
	fmt.Println("[5] :: [See all projects]")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[0] :: [Back]")
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

func InputOption(msgOp string, files []string) int {
	for {
		CreateMenu(files)
		op := InputNumber("[" + msgOp + "] :: ")
		if op < 0 || op > len(files) {
			Info("Option not exist, try again!")
		} else {
			return op
		}
	}
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
