package main

import (
	"errors"
	"fmt"
)

// 1. Error interface is an inbuilt interface that has one method only Error() string
// type error interface {
// 	Error() string
// }

// 2. Custom errors

type UserError struct {
	name string
}

func (u UserError) Error() string {
	return fmt.Sprintf("UserErr: %v has error in data", u.name)
}

//now it can be used as an error

func sendMsg(msg, name string) error {
	if true {
		return UserError{name: name}
	}

	return errors.New("err")
}

func main() {
	fmt.Println(sendMsg("msgg", "sai")) //sai has error in data

}
