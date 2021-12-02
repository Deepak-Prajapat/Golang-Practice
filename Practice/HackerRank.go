package main

import (
	"fmt"
	"strconv"
)

func main() {

	var time string = "11:59:00PM"
	//var time string = "04:59:59AM"
	var militaryTime string = timeConversion(time)
	fmt.Println(militaryTime)
}

func timeConversion(s string) string {

	amOrPm := s[len(s)-2 : len(s)]
	hour := s[0:2]
	var lengthWithoutAMPM = len(s) - 2

	var militaryTime string

	if amOrPm == "PM" {
		if hour == "12" {
			militaryTime = s[0:lengthWithoutAMPM]
		} else if intHour, err := strconv.Atoi(hour); err == nil {
			militaryTime = strconv.Itoa(intHour+12) + s[2:lengthWithoutAMPM]
		}
	} else {
		if hour == "12" {
			militaryTime = "00" + s[2:lengthWithoutAMPM]
		} else {
			militaryTime = s[0:lengthWithoutAMPM]
		}
	}
	return militaryTime
}
