package main

import "github.com/gdamore/tcell/v2"

type Position struct {
	row int32
	col int32
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
	Draw(term tcell.Screen)
	Move(newPos Position)
	SetSize(width int32, height int32)
}
