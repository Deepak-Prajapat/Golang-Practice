package main

import "fmt"

var count int = 0

func main() {
	clearTerminal()
	fmt.Println("count before", count)
	var num float64
	fmt.Print("Enter Number to calculate = ")
	fmt.Scanln(&num)
	var result float64 = fib(num)
	fmt.Println("count", count)
	fmt.Println("result", result)
	fmt.Scanln()
}
