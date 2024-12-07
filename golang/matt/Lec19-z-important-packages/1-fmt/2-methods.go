package main

import (
	"fmt"
	"os"
)

func main() {

	// 1. Append
	// func Append(b []byte, a ...any) []byte
	// func Appendf(b []byte, format string, a ...any) []byte
	// func Appendln(b []byte, a ...any) []byte

	b := []byte("Hello")
	b = fmt.Append(b, " world", 123)
	fmt.Println(string(b)) // Output: Hello world123

	b = []byte("")
	b = fmt.Appendf(b, " %s %d", "world", 123)
	fmt.Println(string(b)) // Output: Hello world 123

	b = []byte("")
	b = fmt.Appendln(b, "world", 123)
	fmt.Println(string(b)) // Output: Hello world 123

	// 2. Errorf  -- func Errorf(format string, a ...any) error
	err := fmt.Errorf("an error occurred: %s", "something went wrong")
	fmt.Println(err.Error()) // Output: an error occurred: something went wrong

	// 3. Printf, Println, Print  you know
	//4. Scan, Scanln, Scanf u know

	// 	func Print(a ...any) (n int, err error)
	// func Printf(format string, a ...any) (n int, err error)
	// func Println(a ...any) (n int, err error)
	// func Scan(a ...any) (n int, err error)
	// func Scanf(format string, a ...any) (n int, err error)
	// func Scanln(a ...any) (n int, err error)

	// 5. Sprintln, Sprintf, Sprint you know... these return formatted strings
	// func Sprint(a ...any) string
	// func Sprintf(format string, a ...any) string
	// func Sprintln(a ...any) string

	// 6. Sscan, Sscanf, Sscanln
	// While simple assignments are fine when you're directly handling well-known values,
	// fmt.Sscan (and similar functions) are designed to provide more control and validation when dealing with dynamic, potentially malformed, or structured data.
	// This makes them essential tools in cases where you need safe, error-resilient parsing of input data.
	// 	func Sscan(str string, a ...any) (n int, err error)
	// func Sscanf(str string, format string, a ...any) (n int, err error)
	// func Sscanln(str string, a ...any) (n int, err error)
	data := "John 30"

	var name string
	var age int

	// Use Sscan to parse the string
	_, err = fmt.Sscan(data, &name, &age)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	_, err = fmt.Sscanf(data, "%s %d", &name, &age)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	_, err = fmt.Sscanln(data, &name, &age)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// 7. Fprint, Fprintf, Fprintln
	// 	func Fprint(w io.Writer, a ...any) (n int, err error)
	// func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	// func Fprintln(w io.Writer, a ...any) (n int, err error)
	_, err = fmt.Fprint(os.Stdout, "Hello, world!")
	_, err = fmt.Fprintf(os.Stdout, "The value is: %d\n", 42)
	_, err = fmt.Fprintln(os.Stdout, "Hello", "world")

	// 8.  Fscan, Fscanln, Fscanf
	// func Fscan(r io.Reader, a ...any) (n int, err error)
	// func Fscanf(r io.Reader, format string, a ...any) (n int, err error)
	// func Fscanln(r io.Reader, a ...any) (n int, err error)
	var a, bb int
	_, err = fmt.Fscan(os.Stdin, &a, &bb)
	_, err = fmt.Fscanf(os.Stdin, "%s %d", &a, &bb)
	_, err = fmt.Fscanln(os.Stdin, &a, &bb)

}
