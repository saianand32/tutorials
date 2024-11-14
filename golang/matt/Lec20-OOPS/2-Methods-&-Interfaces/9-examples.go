package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the source file "a.txt" for reading
	f1, err := os.Open("a.txt") // *os.File implements io.Reader
	if err != nil {
		// Error handling for file opening
		fmt.Println("Error opening source file:", err)
		return
	}
	// Ensure the file is closed when we're done
	defer f1.Close()

	// Create or open the destination file "b.txt" for writing
	f2, err := os.Create("b.txt") // *os.File implements io.Writer
	if err != nil {
		// Error handling for file creation
		fmt.Println("Error creating destination file:", err)
		return
	}
	// Ensure the file is closed when we're done
	defer f2.Close()

	// The io.Copy function expects an io.Writer as the first argument and an io.Reader as the second argument.
	// Both *os.File implement io.Writer and io.Reader interfaces, respectively, so we can use them directly here.
	n, err := io.Copy(f2, f1) // Copies from f1 (io.Reader) to f2 (io.Writer)
	if err != nil {
		// Error handling for the copy operation
		fmt.Println("Error copying data:", err)
		return
	}

	// Output the number of bytes successfully copied
	fmt.Printf("Successfully copied %d bytes\n", n)
}
