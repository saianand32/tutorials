package main

import "fmt"

func main() {
	function := parent()

	function() //6
}

// here you can see that the innerfunction along with its lexical scope gets bundled
func parent() func() {
	a := 2
	return func() {
		fmt.Println(a + 4)
	}
}
