package main

import (
	"fmt"
	"io"
	"os"
)

func writeToFile() {
	content := "this needs to go into file"

	file, err := os.Create("./sai.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	file.Close()
}

func readFromFile() {
	dataByte, err := os.ReadFile("sai.txt")

	if err != nil {
		panic(err)
	}
	fmt.Println(dataByte) //byte form
	fmt.Println(string(dataByte))
}

func main() {
	readFromFile()
}
