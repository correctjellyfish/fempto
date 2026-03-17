package main

// A view of a buffer drawn to the terminal
type BufferView struct {
	// The buffer being viewed
	buffer Buffer
	// The cursor position within the buffer
	cursor Position
	// The position of the BufferView in the terminal
	location Position
	// The number of columns the BufferView occupies
	width int
	// The number of rows the BufferView occupies
	height int
}

// Create a new BufferView from a buffer
func newBufferView(buffer Buffer, location Position, width int, height int) BufferView {
	return BufferView{
		buffer:   buffer,
		location: location,
		width:    width,
		height:   height,
		cursor:   Position{row: 0, col: 0},
	}
}

// Handle an EditCommand
func (bufferView *BufferView) handleCommand(command EditCommand) {
	bufferView.buffer.handleCommand(command)
}
