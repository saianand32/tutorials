package main

import (
	"errors"
	"fmt"
)

func divide(a float64, b float64) (ans float64, err error) { //passing error type as return type

	if b == 0.0 {
		return -1.0, errors.New("Cant divide by 0") //throwing a new error
	}
	return a / b, nil
}

func main() {
	ans, err := divide(2, 0)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ans)
}
