package metacommand

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyCommand = errors.New("got empty command string")
	ErrUnknownCommand = errors.New("unknown command")
	ErrUnknownMetaCommandPrefix = fmt.Errorf("unknown meta command prefix, expected %c", META_COMMAND_PREFIX)
)
