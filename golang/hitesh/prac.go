package main

import (
	"fmt"

	"github.com/badgerodon/collections"
)

func main() {
	stack := collections.NewStack[int]()
	stack.Push(10)
	stack.Push(10)
	stack.Push(10)
	stack.Push(10)

	fmt.Println(stack)
}
