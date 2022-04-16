package sudoku

import (
	"fmt"
)

// sudokusolver is responsible for resolving the first possible solution.

func hCheck(row []rune, toPut rune) bool {
	// checks duplicates horizontally - returns false if match is found
	for _, match := range row {
		if toPut == match {
			return false
		}
	}
	return true
}

func vCheck(rows [][]rune, toPut rune, index int) bool {
	// checks duplicates vertically - returns false if match is found
	for _, row := range rows {
		if row[index] == toPut {
			return false
		}
	}
	return true
}

func squareCheck(rows [][]rune, toPut rune, indexRow int, indexCol int) bool {
	// checks duplicates in 3x3 square grid - returns false if match is found
	var rowToCheck int = (indexRow / 3) * 3 // getting row start
	var colToCheck int = (indexCol / 3) * 3 // getting col start
	for row := rowToCheck; row < rowToCheck+3; row++ {
		for col := colToCheck; col < colToCheck+3; col++ {
			if rows[row][col] == toPut {
				return false
			}
		}
	}
	return true
}

func PutNum(rows [][]rune) bool {
	// fills missing spot in ascending order of grid from [0][0] to [8][8].
	// if not possible to assign a number anymore to empty spot, function assigns spot back to '.', 
	// and moves back to most recent number function was able to assign and tries next possible number. 
	// returns true if can assign a number and calls itself again to find next missing spot.
	for rowIndex, row := range rows {
		for colIndex, _ := range row {
			if rows[rowIndex][colIndex] == '.' {
				for i := '1'; i <= '9'; i++ {
					if hCheck(row, i) && vCheck(rows, i, colIndex) && squareCheck(rows, i, rowIndex, colIndex) {
						rows[rowIndex][colIndex] = i
						if PutNum(rows) {
							return true
						} else {
							rows[rowIndex][colIndex] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func PrintRows(rows [][]rune) {
	// prints sudoku in required format
	for _, row := range rows {
		for index, num := range row {
			fmt.Print(string(num))
			if index != len(row)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
