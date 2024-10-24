package main

import (
	"fmt"
	"io"
	"os"
)

// os.Create("filename"), 
// os.Open("fileName")
// file.Close()

func main() {
	// 1. Create the first file
	file, err := os.Create("sai.txt")
	if err != nil { // Error handling for file creation
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Ensure the first file is closed after use

	// 2. Write to the first file
	_, err = fmt.Fprintln(file, "Hi, this is Saishwar")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Successfully wrote to the file")

	// 3. Create the second file (destination)
	file2, err := os.Create("sai2.txt")
	if err != nil { // Error handling for file creation
		fmt.Println("Error:", err)
		return
	}
	defer file2.Close() // Ensure the second file is closed after use

	// 4. Open the first file in read mode
	file, err = os.Open("sai.txt")
	if err != nil {
		fmt.Println("Error opening sai.txt:", err) // Error handling for file opening
		return
	}

	// 5. Copy the contents from the first file to the second file
	bytesCopied, err := io.Copy(file2, file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error copying file:", err)
		return
	}

	// 6. Report success and the number of bytes copied
	fmt.Printf("Successfully copied %d bytes from sai.txt to sai2.txt\n", bytesCopied)
}
