package main

import (
	"bufio"
	"log"
	"os"
)

// Row represents a single row of text in a Buffer
type Row struct {
	idx     int
	size    int
	content []rune
}

// Create a new row from a string of text, and the
// index of the row in the file
func newRow(text string, rowIndex int) Row {
	r := Row{
		idx:     rowIndex,
		size:    len(text),
		content: []rune(text),
	}

	return r
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
