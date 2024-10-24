# Go Project Structure

-- Go is a modular language, meaning you can create multiple packages for your Go project and split it into modules. This modularity helps in organizing code efficiently, making it more maintainable and scalable.

-- In Go file Names must never start with a . and _ 

## Example: A Simple Go Program

Below is an example of a simple Go program that prints "Hello World" to the console:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```