package crud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"school/pkg/dbcon"
)

type ErrorJson struct {
	Title string `json: title`
	Body  string `json:"body"`
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	var tempStu dbcon.Student
	respError := struct{ Error string }{} // to response error if somethign is not relevant

	//// unmarshal body to student to fetch individual field
	err := json.Unmarshal([]byte(body), &tempStu)
	if err != nil {
		fmt.Println("error", err.Error())
		panic(err)
	}

	if tempStu.Fees < 50000 {
		respError.Error = "Fees should be grater then or equal to 50000"
		json.NewEncoder(w).Encode(respError)
		return
	}

	json.NewDecoder(r.Body).Decode(&tempStu) //Encoding and decoding
	dbcon.Database.Create(&tempStu)
	json.NewEncoder(w).Encode(tempStu)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var students []dbcon.Student

	//// Find all students
	dbcon.Database.Find(&students)
	json.NewEncoder(w).Encode(students) //return in json format while encoding
}

//// New Work for Mocking
type DB interface {
	studentById(stuId string, stu *dbcon.Student) error
}
type db struct{}

var StudentService DB = db{}

func (d db) studentById(stuId string, stu *dbcon.Student) error {
	if err := dbcon.Database.Where("id = ?", stuId).First(&stu).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stuId := r.FormValue("stuId")
	var student dbcon.Student

	err := StudentService.studentById(stuId, &student) //Created mock for this method

	if err != nil {
		var er ErrorJson
		er.Title = err.Error()
		er.Body = "No Student Found For Particular Id = " + stuId
		json.NewEncoder(w).Encode(er)
	} else {
		json.NewEncoder(w).Encode(student)
	}
}
