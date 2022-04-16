package sudoku

func CreateRows(rows [][]rune, args []string) [][]rune {
	for i := 1; i <= 9; i++ {
		rows = append(rows, GetRows(args[i]))
	}
	return rows
}

func GetRows(row string) []rune {
	var getRow []rune
	for _, letter := range row {
		getRow = append(getRow, letter)
	}
	return getRow
}