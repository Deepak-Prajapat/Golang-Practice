package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"unsafe"
)

type Student struct {
	Name   string
	Course string
	Fees   int
}

func main() {
	clearTerminal()
	//// Array of struct in golang
	students := make([]*Student, 0, 10)

	student := new(Student)
	student.Name = "Deepak Prajapati"
	student.Course = "BIT"
	student.Fees = 5000
	students = append(students, student)

	student2 := new(Student)
	student2.Name = "Deepak Prajapati"
	student2.Course = "BIT"
	student2.Fees = 5000

	students = append(students, student2)

	j, _ := json.MarshalIndent(students, "", "    ")
	log.Println(string(j))

	fmt.Println("length of the array", len(students))
	fmt.Println("capacity of the array", cap(students))
	fmt.Println("sixe of students", unsafe.Sizeof(students))
	fmt.Println("sixe of single student", unsafe.Sizeof(student))
}

func clearTerminal() {
	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
