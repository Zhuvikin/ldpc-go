package matrix

type Entry struct {
	Row    int
	Column int

	Left  *Entry
	Right *Entry
	Up    *Entry
	Down  *Entry
}

func (entry Entry) NextInRow() *Entry {
	return entry.Right
}

func (entry Entry) NextInColumn() *Entry {
	return entry.Down
}

func (entry Entry) PreviousInRow() *Entry {
	return entry.Left
}

func (entry Entry) PreviousInColumn() *Entry {
	return entry.Up
}

func (entry Entry) AtEnd() bool {
	return entry.Row < 0
}