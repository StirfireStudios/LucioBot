package main

function runCommandPlay(string *filename) {
	
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
