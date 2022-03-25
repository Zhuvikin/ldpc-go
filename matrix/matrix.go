package matrix

import "fmt"

type Matrix struct {
	ColumnsNumber int
	RowsNumber    int

	Rows    []*Entry
	Columns []*Entry
}

func CreateMatrix(columnsNumber int, rowsNumber int) *Matrix {
	matrix := &Matrix{
		ColumnsNumber: columnsNumber,
		RowsNumber:    rowsNumber,
		Columns:       make([]*Entry, columnsNumber),
		Rows:          make([]*Entry, rowsNumber),
	}

	if rowsNumber <= 0 || columnsNumber <= 0 {
		panic("Invalid number of rows or columns")
	}

	for i := 0; i < rowsNumber; i++ {
		matrix.Rows[i] = &Entry{
			Column: -1,
			Row:    -1,
		}
		var entry = matrix.Rows[i]
		entry.Left, entry.Right, entry.Up, entry.Down = entry, entry, entry, entry
	}

	for j := 0; j < columnsNumber; j++ {
		matrix.Columns[j] = &Entry{
			Column: -1,
			Row:    -1,
		}
		var entry = matrix.Columns[j]
		entry.Left, entry.Right, entry.Up, entry.Down = entry, entry, entry, entry
		entry.Row, entry.Column = -1, -1
	}

	return matrix
}

func (matrix Matrix) Equals(matrix2 *Matrix) bool {
	var e1 *Entry
	var e2 *Entry

	if matrix.RowsNumber != matrix2.RowsNumber || matrix.ColumnsNumber != matrix2.ColumnsNumber {
		return false
	}

	for i := 0; i < matrix.RowsNumber; i++ {
		e1 = matrix.FirstInRow(i)
		e2 = matrix2.FirstInRow(i)
		for !e1.AtEnd() && !e2.AtEnd() {
			if e1.Column != e2.Column {
				return false
			}
			e1 = e1.NextInRow()
			e2 = e2.NextInRow()
		}
		if !e1.AtEnd() || !e2.AtEnd() {
			return false
		}
	}

	return true
}

func (matrix Matrix) Print(name string) {
	fmt.Println(name + ":")
	for i := 0; i < matrix.RowsNumber; i++ {
		for j := 0; j < matrix.ColumnsNumber; j++ {
			entry := matrix.Get(i, j)
			if entry == nil {
				fmt.Print("0")
			} else {
				fmt.Print("1")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (matrix Matrix) FirstInRow(row int) *Entry {
	return matrix.Rows[row].Right
}

func (matrix Matrix) FirstInColumn(column int) *Entry {
	return matrix.Columns[column].Down
}

func (matrix Matrix) LastInRow(row int) *Entry {
	return matrix.Rows[row].Left
}

func (matrix Matrix) LastInColumn(column int) *Entry {
	return matrix.Columns[column].Up
}
