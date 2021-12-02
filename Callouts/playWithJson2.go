package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func clearTerminal() {
	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {
	clearTerminal()
	byteResponse := []byte(getJson())

	var response JsonResponse
	err := json.Unmarshal([]byte(byteResponse), &response)
	if err != nil {
		fmt.Println("error", err.Error())
		panic(err)
	}

	jsonString, _ := json.Marshal(response)
	prettyPrintJson(string(jsonString))

}

type JsonResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Batters struct {
		Batter []struct {
			BatterId   string `json:"id"`
			BatterType string `json:"type"`
		} `json:"batter"`
	} `json:"batters"`
}

func getJson() string {
	return `{
	"id": "0001",
	"type": "donut",
	"name": "Cake",
	"ppu": 0.55,
	"batters":
		{
			"batter":
				[
					{ "id": "1001", "type": "Regular" },
					{ "id": "1002", "type": "Chocolate" },
					{ "id": "1003", "type": "Blueberry" },
					{ "id": "1004", "type": "Devil's Food" }
				]
		},
	"topping":
		[
			{ "id": "5001", "type": "None" },
			{ "id": "5002", "type": "Glazed" },
			{ "id": "5005", "type": "Sugar" },
			{ "id": "5007", "type": "Powdered Sugar" },
			{ "id": "5006", "type": "Chocolate with Sprinkles" },
			{ "id": "5003", "type": "Chocolate" },
			{ "id": "5004", "type": "Maple" }
		]
}`
}

//// This function will help to print json in pretty way
func prettyPrintJson(body string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(body), "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Print(out.String())
}
