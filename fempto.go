package main

import (
	"os"
)

func main() {
	args := os.Args
	var fempto Editor
	if len(args) > 1 {
		fempto = newEditor(&args[1])
	} else {
		fempto = newEditor(nil)
	}
	defer fempto.exit()
	fempto.run()
}
