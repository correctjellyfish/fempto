package main

import "github.com/gdamore/tcell/v2"

type Position struct {
	row int
	col int
}

// Add together two positions
func (pos Position) add(other Position) Position {
	pos.row += other.row
	pos.col += other.col
	return pos
}

// Interface describing the UI Elements of the Editor
type UiElement interface {
	// Draw the element to the provided screen
	draw(term tcell.Screen)
	// Set the position of the element on the screen
	move(newPos Position)
	// Set the size of the element
	setSize(width int, height int)
	commandHandler
}
