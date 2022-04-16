package sudoku

// validation checks if (1) input and Os are valid (2) sudoku is solvable (3) sudoku has unique solution
// returns false if it doesn't meet check requirements

func Solvable(rows [][]rune) bool {
	// checks if sudokusolver is successful
	// returns false if unsuccessful with '.'
	for _, row := range rows {
		for _, dot := range row {
			if dot == '.' {
				return false
			}
		}
	}
	return true
}

func ValidateSudoku(s []string) bool {
	// checks if inputs of Os are valid
	// returns false if inputs not valid
	if len(s) != 10 {
		return false
	}
	for i, str := range s {
		if i == 0 {
			continue
		}
		if len(str) != 9 {
			return false
		}
		for _, c := range str {
			if (c < '0' || c > '9') && (c != '.') {
				return false
			}
		}
	}
	return true
}

func PutNumReverse(rows2 [][]rune) bool {
	// works like PutNum() but in reverse order to check if sudoku has one other solution
	for rowIndex := 8; rowIndex >= 0; rowIndex-- {
		for colIndex := 8; colIndex >= 0; colIndex-- {
			if rows2[rowIndex][colIndex] == '.' { 
				for i := '1'; i <= '9'; i++ { 
					if hCheck(rows2[rowIndex], i) && vCheck(rows2, i, colIndex) && squareCheck(rows2, i, rowIndex, colIndex) {
						rows2[rowIndex][colIndex] = i
						if PutNumReverse(rows2) {
							return true
						} else {
							rows2[rowIndex][colIndex] = '.' 
						}
					}
				}
				return false
			}
		}
	}
	return true
}
