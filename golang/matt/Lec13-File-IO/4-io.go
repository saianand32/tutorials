package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("sai.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, "Hi, this is Saishwar")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Successfully wrote to the file")

	file2, err := os.Create("sai2.txt")
	if err != nil { // Error handling for file creation
		fmt.Println("Error:", err)
		return
	}
	defer file2.Close() // Ensure the second file is closed after use

	file, err = os.Open("sai.txt")
	if err != nil {
		fmt.Println("Error opening sai.txt:", err) // Error handling for file opening
		return
	}

	// ************* IO ****************

	// 1. Copy the contents from the first file to the second file
	bytesCopied, err := io.Copy(file2, file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error copying file:", err)
		return
	}

	fmt.Printf("Successfully copied %d bytes from sai.txt to sai2.txt\n", bytesCopied)

	// ************* io.ReadAll and os.WriteFile ****************

	// 2. Read all contents from sai.txt
	// Re-open sai.txt to read from it again (since the file pointer has already reached the end)
	file, err = os.Open("sai.txt")
	if err != nil {
		fmt.Println("Error opening sai.txt for reading:", err)
		return
	}
	defer file.Close()

	// Using io.ReadAll to read the entire contents of the file
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the content read from the file
	fmt.Println("Contents of sai.txt:")
	fmt.Println(string(content))
}
