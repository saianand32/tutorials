package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// to run the server
// --- go build .
// --- go run main.go

func main() {
	fmt.Println("hello mod")
	//1. install a package --- this creates changes in go.mod and also created a sum.mod file used to verify the package authenticity
	// $ go get -u github.com/gorilla/mux

	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hii All")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>HomePage</h1>"))
}

// Some other go commands cli

// 1. $ go list   //prints myModules
// 2. $ go list all //prints all dependencies even which have not been used in proj
// 3. $ go list -m all //prints all dependencies used in proj
// 3. $ go mod tidy //tidys the dependencies which not used etc
// 3. $ go mod verify //verify the dependencies

// changing go mod file
// 1. go mod edit -go v1.1.7 //changes version of go to 1.1.7
// 1. go mod edit -module myModules //changes module name to myModules

// When u install package by using for example $ go get -u github.com/gorilla/mux
// dependencies installed in cache folder in gopath dir
// but if u want a folder like node_modules use this command
// --- $ go mod vendor
//creates vendor folder
// now to run using pacakges in vendor folder
// ---$ go run -mod=vendor main.go
