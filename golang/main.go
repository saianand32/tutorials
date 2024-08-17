package main

import "fmt"

var WhatIsThe = AnswerToLife()

func AnswerToLife() int { // 1
	return 42
}

func init() { // 2
	WhatIsThe = 0
}

func main() { // 3
	fmt.Println(WhatIsThe)
}
