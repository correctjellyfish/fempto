package main

import (
	"slices"
)

// Detect whitespace
func isWhitespace(char rune) bool {
	switch char {
	case ' ', '-', '_', '\n', '\r':
		return true
	default:
		return false
	}
}

// Line represents a single row of text in a Buffer
type Line struct {
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
func newLine(text string, rowIndex int) Line {
	r := Line{
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
func (line *Line) split(position int) Line {
	// Get the content (this is possibly empty)
	newContent := make([]rune, 0, line.size-position)                                          // Allocate slice that is large enough
	newContent = slices.AppendSeq(newContent, slices.Values(line.content[position:line.size])) // Add the content
	newRow := Line{idx: line.idx, size: len(newContent), content: newContent}                  // Create the new row

	// Shorten the calling row
	line.content = line.content[0:position]
	line.size = len(line.content)

	// Return the newly created row
	return newRow
}

// Join another row to the end of calling row
func (line *Line) join(other Line) {
	line.content = append(line.content, other.content...)
	line.size = len(line.content)
}

// Insert a character into the row at the provided position
func (line *Line) insertChar(char rune, position int) {
	// NOTE: Panics if the position is too large
	line.content = slices.Insert(line.content, position, char)
	line.size++
}

// Delete a character at the specified position in the row.
//
// If the position is too large, does nothing.
func (line *Line) deleteChar(position int) {
	if position >= line.size {
		// Do nothing if position if too large
		return
	}
	// Remove element
	line.content = slices.Delete(line.content, position, position+1)
	line.size--
}

// Return the position of the end of the word which
// position is inside of
func (line *Line) wordEnd(position int) int {
	curIdx := position
	for ; curIdx < line.size; curIdx++ {
		if isWhitespace(line.content[curIdx]) {
			break
		}
	}
	if curIdx >= line.size {
		return line.size - 1
	}
	curIdx--
	return curIdx
}

// Return the position of the start of the word which
// position is inside of
func (line *Line) wordStart(position int) int {
	curIdx := position
	for ; curIdx >= 0; curIdx-- {
		if isWhitespace(line.content[curIdx]) {
			break
		}
	}
	if curIdx < 0 {
		return 0
	}
	curIdx++
	return curIdx
}

// Return the position of the next word following the specified position
func (line *Line) nextWord(position int) int {
	curIdx := position
	for ; curIdx < line.size; curIdx++ {
		if isWhitespace(line.content[curIdx]) {
			break
		}
	}
	if curIdx >= line.size {
		return line.size - 1
	}
	curIdx++
	return curIdx
}

// Return the position of the previous word from the given position
func (line *Line) prevWord(position int) int {
	curIdx := position
	seenWhitespace := false
	for ; curIdx >= 0; curIdx-- {
		if isWhitespace(line.content[curIdx]) {
			if seenWhitespace {
				break
			} else {
				seenWhitespace = true
			}
		}
	}
	if curIdx >= line.size {
		return line.size - 1
	}
	if curIdx < 0 {
		return 0
	}
	curIdx++
	return curIdx
}
