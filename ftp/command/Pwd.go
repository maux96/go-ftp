package commands

type PwdCommand BaseCommand

func (nc PwdCommand) Action(ctx *ConnCtx) error {
	path := ctx.UserPath()
	err := ctx.SendMessage(257, path)
	return err
}
