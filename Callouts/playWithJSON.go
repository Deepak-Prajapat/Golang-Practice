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
	byteArray := []byte(getJson())

	//fmt.Println("", string(byteArray))
	var tmpW TmpWeather
	err := json.Unmarshal([]byte(byteArray), &tmpW)
	if err != nil {
		fmt.Println("error", err.Error())
		panic(err)
	}

	fmt.Println("", tmpW)
	tempList := tmpW.Weather
	tempList2 := tempList[0].Array
	fmt.Println("arraylist", tempList[0].Array)
	fmt.Println("templist2 ", tempList2[0].ArrayName)
	jsonString, _ := json.Marshal(tmpW.Weather[0].Array)
	fmt.Println("jsonString", jsonString)

	prettyPrintJson(string(jsonString))
}

func getJson() string {
	return `{
    "coord": {
        "lon": -0.13,
        "lat": 51.51
    },
    "weather": [
        {
			"id": 300,
			"main": "Drizzle",
			"description": "light intensity drizzle",
			"icon": "09d",
			"array" : [
				{
					"arrayname" : "Raghav"
				},
				{
					"arrayname" : "Deepak"
				}
			]
		}
    ],
    "base": "stations",
    "main": {
        "temp": 280.32,
        "pressure": 1012,
        "humidity": 81,
        "temp_min": 279.15,
        "temp_max": 281.15
    },
    "visibility": 10000,
    "wind": {
        "speed": 4.1,
        "deg": 80
    },
    "clouds": {
        "all": 90
    },
    "dt": 1485789600,
    "sys": {
        "type": 1,
        "id": 5091,
        "message": 0.0103,
        "country": "GB",
        "sunrise": 1485762037,
        "sunset": 1485794875
    },
    "id": 2643743,
    "name": "London",
    "cod": 200
}`
}

/* type Weather struct {
    Location       string
    Weather        string
    Description    string
    Temperature    float32
    MinTemperature float32
    MaxTemperature float32
} */

type TmpWeather struct {
	Location string `json:"name"`
	Weather  []struct {
		ID          uint   `json:id`
		Weather     string `json:"main"`
		Description string `json:"description"`
		Array       []struct {
			ArrayName string `json:"arrayname"`
		} `json:"array"`
	} `json:"weather"`
	Temperature struct {
		Temperature    float32 `json:"temp"`
		MinTemperature float32 `json:"temp_min"`
		MaxTemperature float32 `json:"temp_max"`
	} `json:"main"`
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
