package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	// To clear previous output of terminal
	print("\033[H\033[2J")

	// 5 * 5 matrixOne
	matrixOne := make([][]int, 3)
	for i := range matrixOne {
		matrixOne[i] = make([]int, 3)
	}

	var count int = 1

	for i := 0; i < len(matrixOne); i++ {
		for j := 0; j < len(matrixOne); j++ {
			matrixOne[i][j] = count
			count++
		}
	}

	matrixTwo := make([][]int, 3)
	for i := range matrixTwo {
		matrixTwo[i] = make([]int, 3)
	}

	for i := 0; i < len(matrixTwo); i++ {
		for j := 0; j < len(matrixTwo); j++ {
			matrixTwo[i][j] = count
			count++
		}
	}

	Print(matrixOne)
	fmt.Println("__________--------")
	Print(matrixTwo)
	fmt.Println("__________---------")

	//Print(Addition(matrixOne, num))

	m, err := MatrixMultiplication(matrixOne, matrixTwo)
	if err != nil {
		print("error: ", err.Error())
		return
	}
	Print(m)
}

func MatrixMultiplication(matrixOne [][]int, matrixTwo [][]int) ([][]int, error) {
	//Case where we cannot do multiplication on two matrix
	if len(matrixOne) != len(matrixTwo[0]) {
		return nil, errors.New(`first matrix's columns should be equual to second matrix's row`)
	}

	row := 0
	col := 0
	outputMatrix := make([][]int, 3)
	for i := range outputMatrix {
		outputMatrix[i] = make([]int, 3)
	}
	return MatrixMultiply(matrixOne, matrixTwo, row, col, &outputMatrix)
}
func MatrixMultiply(matrixOne [][]int, matrixTwo [][]int, row int, col int, outputMatrix *[][]int) ([][]int, error) {

	if row == len(matrixTwo) {
		return *outputMatrix, nil
	}

	rowItems := GetRowItems(matrixOne, row)
	colItems := GetColItems(matrixTwo, col)

	var itemOnPosition int = 0
	for i := 0; i < len(colItems); i++ {
		itemOnPosition += rowItems[i] * colItems[i]
	}

	(*outputMatrix)[row][col] = itemOnPosition

	if col >= len(matrixTwo[row])-1 {
		*outputMatrix, _ = MatrixMultiply(matrixOne, matrixTwo, row+1, 0, outputMatrix)
	} else {
		*outputMatrix, _ = MatrixMultiply(matrixOne, matrixTwo, row, col+1, outputMatrix)
	}
	return *outputMatrix, nil
}

func Print(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			print(matrix[i][j], "\t")
		}
		println()
	}
}

func GetRowItems(matrix [][]int, row int) []int {
	var slc []int

	for i := 0; i < len(matrix[row]); i++ {
		slc = append(slc, matrix[row][i])
	}
	return slc
}

func GetColItems(matrix [][]int, col int) []int {
	var slc []int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j == col {
				slc = append(slc, matrix[i][j])
			}
		}
	}
	return slc
}
