package main

import (
	"fmt"
)

func handlePlay(filename string) {
	fmt.Printf("Play: %s", filename)
}

func handleList() {
	fmt.Printf("List")
}

func runCommand(command *messageListener) {
	if command == nil {
		return
	}

	switch command.command {
	case Play:
		handlePlay(command.filename)
	case List:
		handleList()
	}
}
