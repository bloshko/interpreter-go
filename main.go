package main

import (
	"bufio"
	"fmt"
	"os"
)

const ()

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args

	if len(args) < 1 {
		fmt.Println("Not enough arguments")
		os.Exit(64)
	}

	secondArgument := args[1]

	if secondArgument == "--stdio" {
		readAndRunStdio()
	}

	fileContent := readFile(secondArgument)

	fmt.Println(string(fileContent))
}

func readFile(filePath string) []byte {
	fileData, err := os.ReadFile(filePath)
	check(err)

	return fileData

}

func readAndRunStdio() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		check(err)

		run(text)
	}

}

func run(code string) {
	fmt.Println(code)
}
