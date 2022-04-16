package main

import (
	"fmt"
	"os"
	"sudoku"
)


func main() {
	// checks if given argument is a valid sudoku
	if !sudoku.ValidateSudoku(os.Args) {
		fmt.Println("Error") 
		return
	}
	// first solution 
	var rows [][]rune
	rows = sudoku.CreateRows(rows, os.Args)
	sudoku.PutNum(rows)
	if !sudoku.Solvable(rows) {
		fmt.Println("Error")
		return
	}
	// checking if more than one solution
	var rows2 [][]rune
	rows2 = sudoku.CreateRows(rows2, os.Args)
	sudoku.PutNumReverse(rows2)
	for i := 0; i <= 8; i++ {
		for j := 0; j <= 8; j++ {
			if rows[i][j] != rows2[i][j] {
				fmt.Println("Error")
				return
			}
		}
	}
	// prints sudoku
	sudoku.PrintRows(rows)
}
