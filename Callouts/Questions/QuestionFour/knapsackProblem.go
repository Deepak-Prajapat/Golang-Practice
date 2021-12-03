package main

import (
	"fmt"
)

func main() {
	var itemCount int = 4
	var profit [5]int = [5]int{0, 1, 6, 10, 15}
	var weight [5]int = [5]int{0, 4, 1, 3, 5}

	var capacity int = 7

	var table [5][9]int = [5][9]int{}

	for i := 0; i <= itemCount; i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else if weight[i] <= j {
				var values []int = []int{profit[i] + table[i-1][j-weight[i]], table[i-1][j]} //// Formula to find all expected values
				table[i][j] = findMax(values)
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}
	i := itemCount
	j := capacity

	var itemThatWillBePacked []int

	for i > 0 && j > 0 {
		if table[i][j] != table[i-1][j] {
			itemThatWillBePacked = append(itemThatWillBePacked, i)
			j = j - weight[i]
			i--
		} else {
			i--
		}
	}
	fmt.Println("itemThatWillBePacked", itemThatWillBePacked)
}

func findMax(values []int) int {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
