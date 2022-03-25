package matrix

func (matrix *Matrix) Add(matrix2 *Matrix) *Matrix {
	var e1 *Entry
	var e2 *Entry
	r := CreateMatrix(matrix.ColumnsNumber, matrix.RowsNumber)

	if matrix.RowsNumber != r.RowsNumber || matrix.ColumnsNumber != r.ColumnsNumber ||
		matrix2.RowsNumber != r.RowsNumber || matrix2.ColumnsNumber != r.ColumnsNumber {
		panic("Matrices have different dimensions")
	}

	if r == matrix || r == matrix2 {
		panic("Result matrix is the same as one of the operands")
	}

	for i := 0; i < r.RowsNumber; i++ {
		e1 = matrix.FirstInRow(i)
		e2 = matrix2.FirstInRow(i)

		for !e1.AtEnd() && !e2.AtEnd() {
			if e1.Column == e2.Column {
				e1 = e1.NextInRow()
				e2 = e2.NextInRow()
			} else if e1.Column < e2.Column {
				r.Set(e1.Column, i)
				e1 = e1.NextInRow()
			} else {
				r.Set(e2.Column, i)
				e2 = e2.NextInRow()
			}
		}

		for !e1.AtEnd() {
			r.Set(e1.Column, i)
			e1 = e1.NextInRow()
		}

		for !e2.AtEnd() {
			r.Set(e2.Column, i)
			e2 = e2.NextInRow()
		}
	}
	return r
}

func (matrix Matrix) Get(row int, column int) *Entry {
	var re *Entry
	var ce *Entry

	if row < 0 || row >= matrix.RowsNumber || column < 0 || column >= matrix.ColumnsNumber {
		panic("Row or column index out of bounds")
	}

	// Check last entries in row and column
	re = matrix.LastInRow(row)
	if re.AtEnd() || re.Column < column {
		return nil
	}
	if re.Column == column {
		return re
	}

	ce = matrix.LastInColumn(column)
	if ce.AtEnd() || ce.Row < row {
		return nil
	}
	if ce.Row == row {
		return ce
	}

	// Search row and column in parallel, from the front
	re = matrix.FirstInRow(row)
	ce = matrix.FirstInColumn(column)

	for {
		if re.AtEnd() || re.Column > column {
			return nil
		}
		if re.Column == column {
			return re
		}

		if ce.AtEnd() || ce.Row > row {
			return nil
		}
		if ce.Row == row {
			return ce
		}

		re = re.NextInRow()
		ce = ce.NextInColumn()
	}
}

func (matrix Matrix) Set(column int, row int) *Entry {
	var re *Entry
	var columnEntry *Entry
	var newEntry *Entry

	if row < 0 || row >= matrix.RowsNumber || column < 0 || column >= matrix.ColumnsNumber {
		panic("Row or column index out of bounds")
	}

	// Find old entry and return it, or allocate new entry and insert into row.
	re = matrix.LastInRow(row)

	if !re.AtEnd() && re.Column == column {
		return re
	}

	if re.AtEnd() || re.Column < column {
		re = re.Right
	} else {
		re = matrix.FirstInRow(row)
		for {
			if !re.AtEnd() && re.Column == column {
				return re
			}
			if re.AtEnd() || re.Column > column {
				break
			}
			re = re.NextInRow()
		}
	}

	newEntry = &Entry{
		Row:    row,
		Column: column,
		Left:   re.Left,
		Right:  re,
	}

	newEntry.Left.Right = newEntry
	newEntry.Right.Left = newEntry

	// Insert new entry into column.  If we find an existing entry here,
	// the matrix must be garbled, since we didn't find it in the row.
	columnEntry = matrix.LastInColumn(column)

	if !columnEntry.AtEnd() && columnEntry.Row == row {
		panic("Garbled matrix")
	}

	if columnEntry.AtEnd() || columnEntry.Row < row {
		columnEntry = columnEntry.Down
	} else {
		columnEntry = matrix.FirstInColumn(column)
		for {
			if !columnEntry.AtEnd() && columnEntry.Row == row {
				panic("Garbled matrix")
			}
			if columnEntry.AtEnd() || columnEntry.Row > row {
				break
			}
			columnEntry = columnEntry.NextInColumn()
		}
	}

	newEntry.Up = columnEntry.Up
	newEntry.Down = columnEntry
	newEntry.Up.Down = newEntry
	newEntry.Down.Up = newEntry

	return newEntry
}
