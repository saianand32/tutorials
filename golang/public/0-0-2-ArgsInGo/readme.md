# Using Command-Line Arguments in Go

Go provides a simple and efficient way to handle command-line arguments using the `os` package. This allows your program to accept input directly from the command line, which can be useful for various applications like scripts, utilities, or tools that need to process user inputs.

## Example: Handling Command-Line Arguments

The following Go program demonstrates how to handle command-line arguments. It expects a 3 arguments, converts it from a string to an integer, and then prints the integer value:

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Println("Please provide 3 arguments")
		return
	}

	strNum := os.Args[1] // The first argument
    
	// Command-line arguments are always passed as strings
	num, err := strconv.Atoi(strNum) // Convert the string argument to an integer

	if err != nil {
		fmt.Println("couldnt convert to int")
	}

	fmt.Println("Int argument =", num)
	
}
```

## Running the program
1. Using .go file --> ` go run main.go 10 20 30 ` 					--> output: Int arg = 10
2. Build & run    --> `go build` --> `go ./proj_name 10 20 30`      --> output: Int arg = 10




## How the Code Works Checking the Number of Arguments:

The program first checks if an argument was provided using `len(os.Args)`. If fewer than 3 arguments are provided , it prompts the user to provide arguments and exits. 

`os.Args is a slice of strings` where os.Args[0] is the name of the go file or executable, and subsequent elements are the provided command-line arguments. All args recieved are always strings and need to be converted to desired datatype for use. The program uses strconv.Atoi to convert the string argument to an integer.