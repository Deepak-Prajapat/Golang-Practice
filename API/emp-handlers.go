package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

type errorJson struct {
	Title string `json: title`
	Body  string `json:"body"`
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp) //Encoding and decoding
	Database.Create(&emp)                //To Save Data
	json.NewEncoder(w).Encode(emp)       //return in json format while encoding
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	//// Set HEader as json
	w.Header().Set("Content-Type", "application/json")

	var employees []Employee

	//// Find all employees
	Database.Find(&employees)

	//// To Save Data
	json.NewEncoder(w).Encode(employees) //return in json format while encoding

}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {

	fmt.Println("______________________________________________________________________")
	w.Header().Set("Content-Type", "application/json")

	var employee Employee

	if err := Database.First(&employee, mux.Vars(r)["eid"]).Error; err != nil {
		// error handling...
		var indexNo int = strings.LastIndex(r.URL.String(), "/")
		id := r.URL.String()[indexNo+1 : len(r.URL.String())]

		var er errorJson
		er.Body = "No Record Found For Particular Id = " + id
		json.NewEncoder(w).Encode(er)

	} else {
		json.NewEncoder(w).Encode(employee)
	}
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request) {
	var employee Employee

	//// Get String after / in url
	var index int = int(strings.LastIndex(r.URL.String(), "/"))
	var name string = r.URL.String()[index+1 : len(r.URL.String())]

	name, err = url.QueryUnescape(name)

	if err := Database.Where("emp_name = ?", name).First(&employee).Error; err != nil {
		// error handling...
		fmt.Println("error ", err.Error())
		var er errorJson
		er.Body = "Not Found For " + name
		er.Title = err.Error()
		json.NewEncoder(w).Encode(er)
	} else {
		json.NewEncoder(w).Encode(employee)
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("______________________________________________________________________")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("values in r ", *r)
	fmt.Println("method ", r.Method)
	fmt.Println("Get Body", r.GetBody)
	fmt.Println("Post Form", r.PostForm)
	fmt.Println("", r.Response)

	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"]) //Record Fetched
	//// Now we will decode the record
	json.NewDecoder(r.Body).Decode(&employee)

	////Save to database
	Database.Save(&employee)

	////Encode and return
	json.NewEncoder(w).Encode(employee)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var employee Employee

	//// First Way : Fetch And Delete
	/* Database.First(&employee, mux.Vars(r)["eid"])
	Database.Delete(&employee) */

	//// Second Way Direct Delete
	if err := Database.Delete(&employee, mux.Vars(r)["eid"]).Error; err != nil {
		var er errorJson
		er.Body = "Not Found"
		er.Title = err.Error()
		json.NewEncoder(w).Encode(er)
	} else {
		//// Encode Custom message
		json.NewEncoder(w).Encode("Employee Has Been Deleted")
	}

	//// Encoding
	//json.NewEncoder(w).Encode(employee)

}
