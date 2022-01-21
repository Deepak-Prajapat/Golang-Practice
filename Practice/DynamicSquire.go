package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {

	// To clear previous output of terminal
	print("\033[H\033[2J")

	var num int32
	fmt.Print("Enter Number: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		log.Fatalln("error while getting input", err.Error())
	}

	twoDArray := make([][]int, num)

	for i := range twoDArray {
		twoDArray[i] = make([]int, num)
	}

	min := 0
	max := int(num - 1)

	iterateLoop := (num + 1) / 2

	count := 1

	for i := 0; i < int(iterateLoop); i++ {

		//// First Line
		for j := min; j < max+1; j++ {
			twoDArray[i][j] = count
			count++
		}

		for k := min + 1; k < max+1; k++ {
			twoDArray[k][max] = count
			count++
		}

		for j := max - 1; j >= min; j-- {
			twoDArray[max][j] = count
			count++
		}

		for j := max - 1; j > min; j-- {
			twoDArray[j][min] = count
			count++
		}

		min = min + 1
		max = max - 1
	}

	PrintMatrix(num, twoDArray)
}

func PrintMatrix(num int32, twoDArray [][]int) {
	for i := 0; i < int(num); i++ {
		for j := 0; j < int(num); j++ {
			if twoDArray[i][j] < 10 {
				if j == int(num-1) && j < 10 {
					fmt.Print(strconv.Itoa(twoDArray[i][j]) + " ")
					time.Sleep(time.Microsecond * 100)
				} else {
					fmt.Print(strconv.Itoa(twoDArray[i][j]) + " ")
					time.Sleep(time.Microsecond * 100)
				}
			} else {
				fmt.Print(strconv.Itoa(twoDArray[i][j]) + " ")
				time.Sleep(time.Microsecond * 100)
			}
		}
		fmt.Println()
	}
}
