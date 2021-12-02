package main

import (
	"testing"
)

func TestFib(t *testing.T) {
	//// Addingn Test Cases
	testCases := []struct {
		description string
		input       float64
		expected    float64
	}{
		{"successfully calculated fibonachi number ", 5, 5},
		{"successfyyl calculated fibnachi number", 17, 1597},
		{"successfyyl calculated 50 fibnachi number", 50, 12586269025},
		{"successfyyl calculated 0 fibnachi number", 0, 12586269025},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			if result := fib(testCase.input); result != testCase.expected {
				t.Errorf("\nIncorrect value, got: %v, for input:%v, expected: %v", result, testCase.input, testCase.expected)
			}
		})
	}

}
