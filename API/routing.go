package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//// We are creating a router here
func handlerRouting() {
	router := mux.NewRouter()
	router.HandleFunc("/employee", CreateEmployee).Methods("POST")
	router.HandleFunc("/employees", GetEmployees).Methods("GET")
	router.HandleFunc("/employee/{eid}", GetEmployeeByID).Methods("GET")
	router.HandleFunc("/employeeByName/{name}", GetEmployeeByName).Methods("GET")
	router.HandleFunc("/updateEmployee/{eid}", UpdateEmployee).Methods("PUT")    //update employee using id
	router.HandleFunc("/deleteEmployee/{eid}", DeleteEmployee).Methods("DELETE") //update employee using id

	log.Fatal(http.ListenAndServe(":8080", router))
}
