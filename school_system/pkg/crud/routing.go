package crud

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//// setup router
func Routing() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/addstudent", CreateStudent).Methods("POST")
	router.HandleFunc("/students", GetStudents).Methods("GET")
	router.HandleFunc("/studentById", GetStudentById).Methods("GET")
	fmt.Print("Server Started....")
	return router
}
