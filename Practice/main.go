package main

import (
	"fmt"
)

func main() {
	// To clear previous output of terminal
	print("\033[H\033[2J")

	nums := []int{
		0, 11, 10, 1, 3, 2, 6, 5,
	}
	Sort(nums)
	fmt.Println("numbers:", nums)
}

func Sort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j+1] < nums[j] && j < len(nums)-1 {
				nums[j] = nums[j+1] + nums[j]
				nums[j+1] = nums[j] - nums[j+1]
				nums[j] = nums[j] - nums[j+1]
			}
		}
	}
	return nums
}

func ()  {

}