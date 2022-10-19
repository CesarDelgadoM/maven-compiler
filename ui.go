package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Menu() {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                           [MAVEN COMPILER]                           ")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[1] :: [Compile a project]")
	fmt.Println("[2] :: [Menu projects]")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[0] :: [Exit]")
	fmt.Println("------------------------------------------------------------------------")
}

func MenuFilesOptions() {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                           [MENU  PROJECTS]                             ")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[1] :: [Create project]")
	fmt.Println("[2] :: [Open project]")
	fmt.Println("[3] :: [Edit project]")
	fmt.Println("[4] :: [Delete project]")
	fmt.Println("[5] :: [View projects]")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[0] :: [Back]")
	fmt.Println("------------------------------------------------------------------------")
}

func CreateMenu(items []string) {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                          [PROJECTS]                          ")
	fmt.Println("------------------------------------------------------------------------")
	for _, item := range items {
		fmt.Println("->[" + item + "]")
	}
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[0] :: [Back]")
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

func InputOption(msgOp string, keys []string) string {

	CreateMenu(keys)
	op := InputText("[" + msgOp + "] :: ")
	return op
}

func ListNamesProjects() {
	cls()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("                         [PROJECTS]                         ")
	fmt.Println("------------------------------------------------------------------------")
	for key := range projects {
		fmt.Println("->[" + key + "]")
	}
	PrintLine()
	PressEnter()
}

func PressEnter() {
	var enter int
	fmt.Print("prees enter to continue...")
	fmt.Scanln(&enter)
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Info(msg interface{}) {
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[INFO]", msg)
	fmt.Println("------------------------------------------------------------------------")
}

func Error(msg interface{}) {
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("[ERROR]", msg)
	fmt.Println("------------------------------------------------------------------------")
}

func PrintLine() {
	fmt.Println("------------------------------------------------------------------------")
}
