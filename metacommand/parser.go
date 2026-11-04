package metacommand

func Handle(command string) error {

	if len(command) == 0{
		return ErrEmptyCommand
	}

	// confirm command matches prefix
	if command[0] != META_COMMAND_PREFIX {
		return ErrUnknownMetaCommandPrefix
	}

	// strip prefix
	command = command[1:]

	handlerFunc, found := META_COMMANDS[command]
	if !found {
		return ErrUnknownCommand
	}

	return handlerFunc()
}
