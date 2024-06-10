package commands

import (
	"os"
	"strings"
)

type RmdCommand BaseCommand

func (cmd RmdCommand) Action(ctx *ConnCtx) error {
	path := strings.Join(cmd.Args, " ")

	realPath, exists := ctx.GetRealPath(path)
	if !exists {
		err := ctx.SendMessage(501, "Invalid Directory.")
		return err
	}

	err := os.Remove(realPath)
	if err != nil {
		return err
	} else {
		ctx.SendMessage(257, "Directory Removed.")
		return nil
	}
}
