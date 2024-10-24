package main

import (
	"fmt"
	"io"
	"os"
)

// This program reads and prints the contents of one or more text files.
// Usage: go run file_name.go *.txt
// The command line argument "*.txt" will pass all files ending with .txt in the current directory.

func main() {
	// Loop through all command line arguments starting from index 1 (file names).
	for _, fName := range os.Args[1:] {
		// 1. Open the file for reading
		file, err := os.Open(fName)
		if err != nil {
			// Print the error to the standard error (os.Stderr) if the file can't be opened, then continue to the next file.
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// 2. Copy the contents of the file to the standard output (os.Stdout).
		// io.Copy copies data from the file (reader) to the destination (writer), which in this case is Stdout.
		if _, err := io.Copy(os.Stdout, file); err != nil {
			// Print the error to the standard error (os.Stderr) if the copy operation fails, then continue to the next file.
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		// 3. Close the file after successfully reading it.
		file.Close()
	}
}
