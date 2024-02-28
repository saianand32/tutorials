package main

type school struct {
	schName string
}

type Student struct {
	Name   string
	Age    int
	Flag   bool
	Grade  complex128
	school // Embedded struct
}

func main() {
	stud1 := Student{}
	stud1.schName = "ssshss"
}
