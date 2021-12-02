package main

import (
	//"fmt"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Response struct {
	ID      int     `json:ID`
	EmpName string  `json:empname`
	Salary  float64 `json:salary`
	Email   string  `json:email`
}

func (r Response) print() {
	fmt.Println("", r)
}

func main() {

	var responseData []Response = []Response{}

	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	//// Work start from here
	resp, err := http.Get("http://localhost:8080/employees")
	if err != nil {
		fmt.Println("error", err.Error())
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	////converting uncomic json data to struct
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("failed to unmarshal:", err)
	} else {
		fmt.Println("", responseData)
		fmt.Println("=========================")
		prettyPrintJson(string(body))
	}

	var resultArray []map[string]interface{} = []map[string]interface{}{}

	json.Unmarshal([]byte(body), &resultArray)

	fmt.Println("result", resultArray[0])

	jsonString, _ := json.Marshal(resultArray[0])

	//// Convert struct to map
	prettyPrintJson(string(jsonString))

}

func prettyPrintJson(body string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(body), "", "     ")
	if err != nil {
		panic(err)
	}
	fmt.Print(out.String())
}
