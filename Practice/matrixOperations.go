package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	// To clear previous output of terminal
	print("\033[H\033[2J")

	matrixOne, matrixTwo := TakeInput()

	print("\033[H\033[2J")
	fmt.Println("First Matrix\n````````````````")
	Print(matrixOne)
	fmt.Println("\nSecond Matrix\n`````````````````")
	Print(matrixTwo)

	//Print(Addition(matrixOne, num))

	m, err := MatrixMultiplication(matrixOne, matrixTwo)
	if err != nil {
		print("\nSigma Rule #1001 ", err.Error())
		print("\n`````````````````\n")
		return
	}
	fmt.Println("\nMultiplication Result\n```````````````````````")
	Print(m)
}

func MatrixMultiplication(matrixOne [][]int, matrixTwo [][]int) ([][]int, error) {
	//Case where we cannot do multiplication on two matrix
	if len(matrixOne[0]) != len(matrixTwo) {
		return nil, errors.New(`first matrix's columns should be equal to second matrix's row`)
	}

	row := 0
	col := 0
	outputMatrix := make([][]int, len(matrixOne))
	for i := range outputMatrix {
		outputMatrix[i] = make([]int, len(matrixTwo[0]))
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
		for j := 0; j < len(matrix[0]); j++ {
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

func TakeInput() ([][]int, [][]int) {
	fmt.Println("Enter Rows and Columns For First Matrix")
	fmt.Println("``````````````````````````````````````````")
	var rowCount int
	fmt.Print("How many rows you want in first matrix = ")
	fmt.Scanln(&rowCount)

	var colCount int
	fmt.Print("How many columns you want in first matrix = ")
	fmt.Scanln(&colCount)

	matrixOne := make([][]int, rowCount)
	for i := range matrixOne {
		matrixOne[i] = make([]int, colCount)
	}

	for i := 1; i <= rowCount; i++ {
		for j := 1; j <= colCount; j++ {
			fmt.Print("Enter item for position [", i, ",", j, "] = ")
			fmt.Scanln(&matrixOne[i-1][j-1])
		}
	}

	//For Second Matrix
	fmt.Println("\n\nEnter Rows and Columns For Second Matrix")
	fmt.Println("`````````````````````````````````````````")
	fmt.Print("How many rows you want in Second matrix = ")
	fmt.Scanln(&rowCount)

	fmt.Print("How many columns you want in Second matrix = ")
	fmt.Scanln(&colCount)

	matrixTwo := make([][]int, rowCount)
	for i := range matrixTwo {
		matrixTwo[i] = make([]int, colCount)
	}

	for i := 1; i <= rowCount; i++ {
		for j := 1; j <= colCount; j++ {
			fmt.Print("Enter item for position [", i, ",", j, "] = ")
			fmt.Scanln(&matrixTwo[i-1][j-1])
		}
	}

	return matrixOne, matrixTwo
}
