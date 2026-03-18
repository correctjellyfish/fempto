package main

import "log"

// NOTE: The visitor pattern here is mostly to allow for
// some UI elements to just pass the command through

// Commands sent from the editor
type command interface {
	// Visit a handler (just call )
	visit(handler commandHandler)
}

// Type which can handle commands
type commandHandler interface {
	// Handle the passed command
	acceptCommand(command command)
}

// Insert Command
type insertCommand struct {
	char rune
}

func (insert insertCommand) visit(handler commandHandler) {
	switch h := handler.(type) {
	case *BufferView:
		edit := insertEdit{
			char:     insert.char,
			position: h.cursor,
		}
		h.buffer.handleEdit(edit)
	default:
		log.Fatalf("Insert Command called %T instead of BufferView", h)
	}
}

// end Insert Command

// Quit Command
// Quit the editor
type quitCommand struct{}

// No-op, the editor will handle this
func (quit quitCommand) visit(handler commandHandler) {
	switch h := handler.(type) {
	case *Editor:
		h.shouldQuit = true
	default:
		log.Fatalf("Quit command called on %T instead of editor", h)
	}
}

// end Quit Command

// Resize Command
type resizeCommand struct {
	width  int
	height int
}

func (resize resizeCommand) visit(handler commandHandler) {
	switch h := handler.(type) {
	case *Editor:
		h.view.setSize(resize.width, resize.height)
	default:
		log.Fatalf("Resize command called on %T instead of editor", h)
	}
}

// end Resize Command
