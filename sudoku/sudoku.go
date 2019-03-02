package main

import (
	"fmt"
)

func main() {
	input := [][]int{
		{5, 1, 7, 6, 0, 0, 0, 3, 4},
		{2, 8, 9, 0, 0, 4, 0, 0, 0},
		{3, 4, 6, 2, 0, 5, 0, 9, 0},
		{6, 0, 2, 0, 0, 0, 0, 1, 0},
		{0, 3, 8, 0, 0, 6, 0, 4, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 8},
		{7, 0, 3, 4, 0, 0, 5, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	valid, backtrackCounter := solveSudoku(input)
	fmt.Printf("Valid %v, backtracks %v \n", valid, backtrackCounter)

	input2 := [][]int{
		{5, 1, 7, 6, 0, 0, 0, 3, 4},
		{0, 8, 9, 0, 0, 4, 0, 0, 0},
		{3, 0, 6, 2, 0, 5, 0, 9, 0},
		{6, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 3, 0, 0, 0, 6, 0, 4, 7},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 9, 0, 0, 0, 0, 0, 7, 8},
		{7, 0, 3, 4, 0, 0, 5, 6, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	valid1, backtrackCounter := solveSudoku(input2)
	fmt.Printf("Valid %v, backtracks %v \n", valid1, backtrackCounter)

	inpd := [][]int{
		{1, 0, 5, 7, 0, 2, 6, 3, 8},
		{2, 0, 0, 0, 0, 6, 0, 0, 5},
		{0, 6, 3, 8, 4, 0, 2, 1, 0},
		{0, 5, 9, 2, 0, 1, 3, 8, 0},
		{0, 0, 2, 0, 5, 8, 0, 0, 9},
		{7, 1, 0, 0, 3, 0, 5, 0, 2},
		{0, 0, 4, 5, 6, 0, 7, 2, 0},
		{5, 0, 0, 0, 0, 4, 0, 6, 3},
		{3, 2, 6, 1, 0, 7, 0, 0, 4},
	}

	valid2, backtrackCounter := solveSudoku(inpd)
	fmt.Printf("Valid %v, backtracks %v \n", valid2, backtrackCounter)

	hard := [][]int{
		{8, 5, 0, 0, 0, 2, 4, 0, 0},
		{7, 2, 0, 0, 0, 0, 0, 0, 9},
		{0, 0, 4, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 7, 0, 0, 2},
		{3, 0, 5, 0, 0, 0, 9, 0, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 8, 0, 0, 7, 0},
		{0, 1, 7, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 3, 6, 0, 4, 0},
	}

	valid3, backtrackCounter := solveSudoku(hard)
	fmt.Printf("Valid %v, backtracks %v \n", valid3, backtrackCounter)

	diff := [][]int{
		{0, 0, 5, 3, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 0, 2, 0},
		{0, 7, 0, 0, 1, 0, 5, 0, 0},
		{4, 0, 0, 0, 0, 5, 3, 0, 0},
		{0, 1, 0, 0, 7, 0, 0, 0, 6},
		{0, 0, 3, 2, 0, 0, 0, 8, 0},
		{0, 6, 0, 5, 0, 0, 0, 0, 9},
		{0, 0, 4, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 9, 7, 0, 0},
	}

	valid4, backtrackCounter := solveSudoku(diff)
	fmt.Printf("Valid %v, backtracks %v \n", valid4, backtrackCounter)
}

func solveSudoku(numbers [][]int) (bool, int) {
	backtracks := 0
	i, j := findNextCellToFill(numbers)
	if i == -1 {
		return true, backtracks
	}

	for n := 1; n <= 9; n++ {
		if isValid(numbers, i, j, n) {
			numbers[i][j] = n

			valid, backtrackCounter := solveSudoku(numbers)
			backtracks += backtrackCounter
			if valid {
				return true, backtracks
			}
			backtracks++
			numbers[i][j] = 0
		}
	}
	return false, backtracks
}

func findNextCellToFill(numbers [][]int) (int, int) {
	for i, row := range numbers {
		for j, _ := range row {
			if numbers[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isValid(numbers [][]int, rowIndex int, columnIndex int, e int) bool {
	rowOk := !containsInRow(numbers[rowIndex], e)
	if rowOk {
		columnOk := !containsInColumn(numbers, columnIndex, e)
		if columnOk {
			sectorOk := !containsInSector(numbers, rowIndex, columnIndex, e)
			if sectorOk {
				return true
			}
		}
	}
	return false
}

func containsInRow(row []int, number int) bool {
	for _, n := range row {
		if n == number {
			return true
		}
	}
	return false
}

func containsInColumn(numbers [][]int, columnIndex int, number int) bool {
	for i := 0; i <= 8; i++ {
		if numbers[i][columnIndex] == number {
			return true
		}
	}
	return false
}

func containsInSector(numbers [][]int, rowIndex int, columnIndex int, number int) bool {
	x, y := getTopLeftCellOfSector(rowIndex, columnIndex)
	for i := x; i <= x+2; i++ {
		for j := y; j <= y+2; j++ {
			if numbers[i][j] == number {
				return true
			}
		}
	}
	return false
}

func getTopLeftCellOfSector(rowIndex int, columnIndex int) (int, int) {
	return rowIndex / 3 * 3, columnIndex / 3 * 3
}
