package main

import (
	"fmt"
	"strconv"
)

func main() {
	// To clear previous output of terminal
	print("\033[H\033[2J")

	var time string = "12:59:00AM"
	fmt.Println("Previous Time:", time)
	//var time string = "04:59:59AM"
	var militaryTime string = timeConversion(time)
	fmt.Println("In militry format :", militaryTime)
	fmt.Println("chlo ise to yha khatam krte h:")
}

/*
 * It will convert 12Hour format time to 24 hour
 */
func timeConversion(timeString string) string {
	amOrPm := timeString[len(timeString)-2 : len(timeString)]
	hour := timeString[0:2]
	var lengthWithoutAMPM = len(timeString) - 2

	var militaryTime string

	if amOrPm == "PM" {
		if hour == "12" {
			militaryTime = timeString[0:lengthWithoutAMPM]
		} else if intHour, err := strconv.Atoi(hour); err == nil {
			militaryTime = strconv.Itoa(intHour+12) + timeString[2:lengthWithoutAMPM]
		}
	} else if hour == "12" {
		militaryTime = "00" + timeString[2:lengthWithoutAMPM]
	} else {
		militaryTime = timeString[0:lengthWithoutAMPM]
	}
	return militaryTime
}
