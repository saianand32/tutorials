package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// 1. taking stdinput
	reader := bufio.NewReader(os.Stdin)

	data, err := reader.ReadString('\n')

	if err == nil {
		fmt.Println(data)
	}

	// 2. reading file line by line (implement word count program)

	for _, fName := range os.Args[1:] {

		lineCount, wordCount, charCount := 0, 0, 0

		file, err := os.Open(fName)
		if err != nil {
			fmt.Println(err)
			continue
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			lineCount++
			line := scanner.Text()
			charCount += len(line)
			wordCount += len(strings.Split(line, " "))
		}

		fmt.Println("lc: ", lineCount)
		fmt.Println("wc: ", wordCount)
		fmt.Println("cc: ", charCount)
	}

}
