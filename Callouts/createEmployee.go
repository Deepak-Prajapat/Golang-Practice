package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	//// Employee Name
	/* fmt.Print("Enter Employee Name = ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = input[0 : len(input)-2]

	//// Employee Salary
	fmt.Print("Enter Salary = ")
	var salary float64
	fmt.Scanln(&salary)

	fmt.Print("Enter Email = ")
	var email string
	fmt.Scanln(&email) */

	employeeMap := map[string]interface{}{
		"empname": "Raghav Prajapti",
		"salary":  1000,
		"email":   "email",
	}

	fmt.Println("length", len(employeeMap))

	json_data, err := json.Marshal(employeeMap)

	if err != nil {
		log.Fatal(err)
	}

	///// send post
	resp, err := http.Post("http://localhost:8080/employee", "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	prettyPrintJson(string(body))
}

func prettyPrintJson(body string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(body), "", "     ")
	if err != nil {
		panic(err)
	}
	fmt.Print(out.String())
}

