package main

// Commands sent from the editor to the Buffer
type EditCommand interface {
	execute(buffer Buffer)
	undo(buffer Buffer)
}
