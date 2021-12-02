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

/* type SendBody struct {
	ID      int     `json:id`
	EmpName string  `json:empname`
	Salary  float64 `json:salary`
	Email   string  `json:email`
} */

func main() {
	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	//// Work start from here
	//// name := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Employee Name To Fetch = ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = input[0 : len(input)-2]

	//name2 := string(input)[0 : len(string(input))-2]

	url := "http://localhost:8080/employeeByName/" + input
	fmt.Println("url", url)
	fmt.Println("name input", string(input))
	//// Work start from here
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error", err.Error())
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	prettyPrintJson(string(body))
	fmt.Println("-----------------------------------------")

	// Declared an empty map interface
	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)

	fmt.Println("result", result)

	fmt.Println("emp name", result["empname"])
}

func prettyPrintJson(body string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(body), "", "     ")
	if err != nil {
		panic(err)
	}
	fmt.Print(out.String())
}
