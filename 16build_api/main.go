package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file

type Course struct {
	CourseID   string  `json:"course_id"`
	CourseName string  `json:"course_name"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
	Email    string `json:"email"`
}

// fake DB
var courses []Course

// middleware, helper functions - file
func (c *Course) IsEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - Learning golang")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseID: "1", CourseName: "ReactJS Bootcamp", Price: 299, Author: &Author{FullName: "John Doe", Website: "johndoe.com", Email: "john.doe@example.com"}})
	courses = append(courses, Course{CourseID: "2", CourseName: "Angular Bootcamp", Price: 199, Author: &Author{FullName: "Jane Smith", Website: "janesmith.com", Email: "jane.smith@example.com"}})
	courses = append(courses, Course{CourseID: "3", CourseName: "VueJS Bootcamp", Price: 299, Author: &Author{FullName: "Alice Johnson", Website: "alicejohnson.com", Email: "alice.johnson@example.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	w.Header().Set("Content-Type", "application/json")

	/// grab id from request
	params := mux.Vars(r)

	fmt.Println("Params are:", params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}

	// generate unique id, string
	course.CourseID = fmt.Sprintf("COURSE%v", len(courses)+1)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop, find matching id and remove the course from slice
	for index, course := range courses {
		if course.CourseID == params["id"] {
			// remove the course
			courses = append(courses[:index], courses[index+1:]...)
			var cases Course
			_ = json.NewDecoder(r.Body).Decode(&cases)
			cases.CourseID = params["id"]
			courses = append(courses, cases)
			json.NewEncoder(w).Encode(cases)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)
	// loop, find matching id and remove the course from slice
	for index, course := range courses {
		if course.CourseID == params["id"] {
			// remove the course
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The course has been deleted successfully")
			break
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}
