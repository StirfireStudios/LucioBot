package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func handlePlay(filename string) {
	fmt.Printf("Play: %s", filename)
}

func handleList() {
	audioDir := Config().AudioDir
	entries, err := ioutil.ReadDir(audioDir)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), "mp3") {
			fmt.Printf("Found MP3 file!")
		}
	}
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
