package main

import "github.com/gdamore/tcell/v2"

// A view of a buffer drawn to the terminal
type BufferView struct {
	// The buffer being viewed
	buffer Buffer
	// The cursor position within the buffer
	cursor Position
	// The position of the BufferView in the terminal
	position Position
	// The number of columns the BufferView occupies
	width int
	// The number of rows the BufferView occupies
	height int
	// Whether the buffer view needs to be redrawn
	needsRedraw bool
}

// Create a new BufferView from a buffer
func newBufferView(buffer Buffer, position Position, width int, height int) BufferView {
	return BufferView{
		// The buffer being drawn to the screen
		buffer: buffer,
		// Location of the bufferview in the terminal
		position: position,
		// The width of the view
		width: width,
		// The height of the view
		height: height,
		// Position of the cursor within the
		cursor: Position{row: 0, col: 0},
	}
}

// Draw the buffer to the screen
func (bufferView *BufferView) draw(term tcell.Screen) {
	if !bufferView.needsRedraw {
		return
	}
}

// Move the bufferView to a new location
func (bufferView *BufferView) move(newPos Position) {
	bufferView.position = newPos
}

// Set the size of the bufferView
func (bufferView *BufferView) setSize(width int, height int) {
	bufferView.width = width
	bufferView.height = height
	bufferView.needsRedraw = true
}

// Handle an EditCommand
func (bufferView *BufferView) acceptCommand(command command) {
	command.visit(bufferView)
}
