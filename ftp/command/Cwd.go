package commands

import "strings"

type CwdCommand BaseCommand

func (cmd CwdCommand) Action(ctx *ConnCtx) error {
	newPath := strings.Join(cmd.Args, " ")
	err := ctx.ChangePath(newPath)
	if err != nil {
		ctx.SendMessage(501, "Invalid Directory.")
		return err
	} else {
		ctx.SendMessage(250, "Directory Changed.")
		return nil
	}
}
