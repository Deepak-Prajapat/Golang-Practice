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


	// initialize http client
    client := &http.Client{}

	var id string
	fmt.Print("Enter Employee ID to Fetch = ")
	fmt.Scanln(&id)

	///// send post
	 // set the HTTP method, url, and request body
    req, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/deleteEmployee/" + id, nil)
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
