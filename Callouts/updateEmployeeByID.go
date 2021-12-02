package main

import (
	"bufio"
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
	fmt.Print("Enter Employee Name = ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = input[0 : len(input)-2]

	//// Employee Salary
	fmt.Print("Enter Salary = ")
	var salary float64
	fmt.Scanln(&salary)

	fmt.Print("Enter Email = ")
	var email string
	fmt.Scanln(&email)

	employeeMap := map[string]interface{}{
		"empname": input,
		"salary":  salary,
		"email":   email,
	}

	// initialize http client
    client := &http.Client{}

	json_data, err := json.Marshal(employeeMap)

	if err != nil {
		log.Fatal(err)
	}

	///// send post
	 // set the HTTP method, url, and request body
    req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/updateEmployee/10", bytes.NewBuffer(json_data))
    if err != nil {
        panic(err)
    }

	// set the request header Content-Type for json
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
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
