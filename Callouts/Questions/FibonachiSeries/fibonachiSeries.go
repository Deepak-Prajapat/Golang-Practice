package main

import (
	"os"
	"os/exec"
)

func clearTerminal() {
	//// This will clear terminal's previous output
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

///// Fibonacci number with dynamic programming
func fib(num float64) float64 {

	//// To Store Calculated Data From findFib anonymous Function
	calculatedResults := make(map[float64]float64)

	/* Recursive Anonymous Method That Will Return Fibonacci Number */
	var findFib func(float64) float64
	findFib = func(num float64) float64 {
		count++
		if num <= 1 {
			return num
		}

		if _, ok := calculatedResults[num]; ok {
			return calculatedResults[num]
		}

		calculatedResults[num] = findFib(num-1) + findFib(num-2)
		return calculatedResults[num]
	}

	return findFib(num)
}
