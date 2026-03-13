package main

import (
	"bufio"
	"log"
	"os"
	"slices"
)

// Constant for whitespace
var whitespace []rune = []rune{' ', '-', '_'}

// Row represents a single row of text in a Buffer
type Row struct {
	// The index of the row within the buffer,
	// mostly used for accessing neighboring rows
	idx int
	// Length of the content in the Row
	size int
	// The text in the Row
	content []rune
}

// Create a new row from a string of text, and the
// index of the row in the file
func newRow(text string, rowIndex int) Row {
	r := Row{
		idx:     rowIndex,
		content: []rune(text),
	}
	r.size = len(r.content)

	return r
}

// Split a row starting at the given position, shortening the
// calling row and returning the remainder of the line
// as a new Row. This new row will share its idx with the
// calling row, and so that needs to be updated.
func (row *Row) split(position int) Row {
	// Get the content (this is possibly empty)
	newContent := make([]rune, 0, row.size-position)                                         // Allocate slice that is large enough
	newContent = slices.AppendSeq(newContent, slices.Values(row.content[position:row.size])) // Add the content
	newRow := Row{idx: row.idx, size: len(newContent), content: newContent}                  // Create the new row

	// Shorten the calling row
	row.content = row.content[0:position]
	row.size = len(row.content)

	// Return the newly created row
	return newRow
}

// Join another row to the end of calling row
func (row *Row) join(other Row) {
	row.content = append(row.content, other.content...)
	row.size = len(row.content)
}

// Insert a character into the row at the provided position
func (row *Row) insertChar(position int, char rune) {
	// NOTE: Panics if the position is too large
	row.content = slices.Insert(row.content, position, char)
	row.size++
}

// Delete a character at the specified position in the row.
//
// If the position is too large, does nothing.
func (row *Row) deleteChar(position int) {
	if position >= row.size {
		// Do nothing if position if too large
		return
	}
	// Remove element
	row.content = slices.Delete(row.content, position, position+1)
	row.size--
}

// Return the position of the next word following the
func (row *Row) nextWord(position int) int {
	if position >= row.size {
		return row.size - 1
	}
	curIdx := position
	for ; curIdx < row.size; curIdx++ {
		if slices.Contains(whitespace, row.content[curIdx]) {
			break
		}
	}
	return curIdx
}

// Buffer is a representation of lines of text, normally read from a file
type Buffer struct {
	nrows int
	rows  []Row
}

// Create a new Buffer from a text file. If path is an empty string,
// creates an empty buffer.
func newBuffer(path string) Buffer {
	if path == "" { // Empty strings are treated as nil
		return Buffer{
			nrows: 0,
			rows:  make([]Row, 0),
		}
	}
	// Read the file
	rows := make([]Row, 0)
	rowIdx := 0
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read each line, incrementing the row index and
	// adding the line to the Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, newRow(scanner.Text(), rowIdx))
		rowIdx++
	}
	return Buffer{
		nrows: rowIdx,
		rows:  rows,
	}
}
