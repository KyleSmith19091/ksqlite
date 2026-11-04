package metacommand

import (
	"os"
)

const META_COMMAND_PREFIX = '.'

var META_COMMANDS = map[string]func() error{
	"exit": func() error {
		os.Exit(0)
		return nil
	},
}
