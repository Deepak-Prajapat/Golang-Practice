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
	"strconv"
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

	var id int = 6
	//// Work start from here
	resp, err := http.Get("http://localhost:8080/employee/" + strconv.Itoa(id))
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
}

func prettyPrintJson(body string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(body), "", "     ")
	if err != nil {
		panic(err)
	}
	fmt.Print(out.String())
}
