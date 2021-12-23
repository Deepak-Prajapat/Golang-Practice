package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	ID       int    `json:"ID,omitempty"`
	Name     string `json:"name,omitempty"`
	Course   string `json:"course,omitempty"`
	Fees     int32  `json:"fees,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Subjects string `json:"subjects,omitempty"`
}

func main() {
	//// This init will clear terminal's previous output on before main loads
	fmt.Print("\033[H\033[2J")

	//// Work start from here
	stu := studentServiceimps{}
	resp, err := stu.GetStudentById()
	if err != nil {
		fmt.Println("error >>> ", err)
	}

	fmt.Println("response >>> ", *resp)

}

///// For Mocking
var StudentService studentService = studentServiceimps{}

type studentService interface {
	GetStudentById() (*Student, error)
}

type studentServiceimps struct{}

func (service studentServiceimps) GetStudentById() (*Student, error) {

	var stuId int = 125
	resp, err := http.Get("http://localhost:8080" + "/studentById?stuId=" + strconv.Itoa(stuId))

	if err != nil {
		fmt.Println("error", err.Error())
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var student Student
	err = json.Unmarshal([]byte(body), &student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

var (
	ResultController = resultCont{}
)

type resultCont struct{}

func (controller resultCont) GetResult() (*Student, error) {
	stu, err := StudentService.GetStudentById()
	if err != nil {
		return nil, err
	}
	return stu, nil
}
