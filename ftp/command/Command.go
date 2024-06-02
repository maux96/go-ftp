package commands

import (
	"strings"
)

type Command interface {
	Action(ctx *ConnCtx) error
}

type RawCommand string
type BaseCommand struct {
	CommandName string
	Args        []string
}

func (rc RawCommand) GetCommand() BaseCommand {
	tokens := strings.Fields(string(rc))
	return BaseCommand{
		CommandName: tokens[0],
		Args:        tokens[1:],
	}
}
