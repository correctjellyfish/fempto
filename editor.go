package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

// Struct implementing the main editor interface
type Editor struct {
	// The root UiElement
	view UiElement
	// The terminal screen where the editor is being drawn
	screen tcell.Screen
	// The current width of the editor
	width int
	// The current height of the editor
	height int
	// Whether the editor should quit
	shouldQuit bool
}

func newEditor(path *string) Editor {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Panicf("%+v", err)
	}

	w, h := s.Size()

	buffer := newBuffer(path)
	bufferView := newBufferView(buffer, Position{0, 0}, w, h)

	return Editor{
		view:       &bufferView,
		screen:     s,
		width:      w,
		height:     h,
		shouldQuit: false,
	}
}

// Run the editor
func (editor *Editor) run() {
	for !editor.shouldQuit {
		editor.screen.Clear()
		command := editor.handleInput()
		if command != nil {
			editor.acceptCommand(command)
		}

	}
}

// Exit the editor
func (editor *Editor) exit() {
	// Clean up the screen
	editor.screen.Fini()
}

func (editor *Editor) handleInput() command {
	ev := editor.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			return &quitCommand{}
		}
	}
	return nil
}

// Accept a command
func (editor *Editor) acceptCommand(command command) {
	switch c := command.(type) {
	// Commands that act on the editor
	case *quitCommand, *resizeCommand:
		c.visit(editor)
	// Commands that act on other parts of the UI
	default:
		editor.view.acceptCommand(command)
	}
}
