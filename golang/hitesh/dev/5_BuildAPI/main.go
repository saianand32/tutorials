package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"slices"
	"strconv"

	"github.com/gorilla/mux"
)

// Models for course
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fakeDB
var courses []Course

// middleware
func (c *Course) isEmpty() bool {
	if c.CourseName == "" {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("API Learn")
	router := mux.NewRouter()

	//seeding
	courses = slices.Insert(courses, len(courses), Course{
		CourseId:    "2",
		CourseName:  "ReactJs",
		CoursePrice: 200,
		Author: &Author{
			Fullname: "Hitesh",
			Website:  "www.google.com",
		},
	})

	courses = slices.Insert(courses, len(courses), Course{
		CourseId:    "4",
		CourseName:  "Angular",
		CoursePrice: 300,
		Author: &Author{
			Fullname: "Hitesh",
			Website:  "www.fb.com",
		},
	})

	//Routing

	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//Listen on port
	log.Fatal(http.ListenAndServe(":4000", router))
}

//controllers

// 1. serveHome

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Courses</h1>"))
}

// 2. getAll courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// 3. getOneCourse
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//grab id from req
	params := mux.Vars(r)

	//find matching id and find resp

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with CourseId = " + string(params["id"]))
}

// 4. create 1 course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("Please send data")
		return
	}

	//generate unique id (string)
	course.CourseId = strconv.Itoa(rand.Intn(2000))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)

}

//5. Update one course

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a response when id is not found
}

// 6. Delete one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}
