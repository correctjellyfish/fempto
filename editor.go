package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

// Struct implementing the main editor interface
type Editor struct {
	// The Buffer currently being viewed
	view BufferView
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
		view:       bufferView,
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
			editor.view.handleCommand(*command)
		}

	}

}

// Exit the editor
func (editor *Editor) exit() {
	// Clean up the screen
	editor.screen.Fini()
}

func (editor *Editor) handleInput() *EditCommand {
	ev := editor.screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
			editor.shouldQuit = true
		}

	}
	return nil
}
