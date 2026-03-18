package main

// An edit to a buffer
type edit interface {
	// Execute the edit
	execute()
	// Undo the edit
	undo()
}

// start insertEdit
type insertEdit struct {
	char     rune
	position Position
}

// Insert a character
func (insertEdit insertEdit) execute() {
}

// Undo a character insert
func (insertEdit insertEdit) undo() {
}

// end insertEdit
