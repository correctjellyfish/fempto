package main

import (
	"bufio"
	"errors"
	"log"
	"os"
)

// Buffer is a representation of lines of text, normally read from a file
type Buffer struct {
	// The number of rows in the buffer
	nlines  int
	lines   []Line
	file    *string
	history []edit
}

// Create a new Buffer from a text file. If path is an empty string,
// creates an empty buffer.
func newBuffer(path *string) Buffer {
	if path == nil { // Empty strings are treated as nil
		return Buffer{
			nlines:  0,
			lines:   make([]Line, 0),
			history: make([]edit, 0),
			file:    nil,
		}
	}
	// Read the file
	rows := make([]Line, 0)
	rowIdx := 0
	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read each line, incrementing the row index and
	// adding the line to the Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, newLine(scanner.Text(), rowIdx))
		rowIdx++
	}
	return Buffer{
		nlines:  rowIdx,
		lines:   rows,
		history: make([]edit, 0),
		file:    path,
	}
}

// Write the contents of the buffer to the associated file
func (buffer *Buffer) write() error {
	if buffer.file == nil {
		log.Print("Tried to write when no file associated with buffer")
		return NoFileError
	}
	file, err := os.Create(*buffer.file)
	if err != nil {
		log.Print("Failed to open file")
		return FileOpenError
	}
	defer file.Close()

	for idx, row := range buffer.lines {
		_, err := file.WriteString(string(append(row.content, '\n')))
		if err != nil {
			log.Printf("Failed to write line %v to file %v, line contents: %v", idx, buffer.file, string(row.content))
			return WriteError
		}
	}

	return nil
}

// Insert a character at the specified position
func (buffer *Buffer) insertChar(char rune, position Position) {
	buffer.lines[position.row].insertChar(char, position.col)
}

// Delete a character at the specified position
func (buffer *Buffer) deleteChar(position Position) {
	buffer.lines[position.row].deleteChar(position.col)
}

// Handle the execution of an EditCommand
func (buffer *Buffer) handleEdit(edit edit) {
	edit.execute()
	buffer.history = append(buffer.history, edit)
}

// Errors
var (
	FileOpenError = errors.New("Failed to open file")
	WriteError    = errors.New("Failed to write to file")
	NoFileError   = errors.New("No file to write to")
)
