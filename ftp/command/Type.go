package commands

type TypeCommand BaseCommand

func (cmd TypeCommand) Action(ctx *ConnCtx) error {
	err := ctx.SendMessage(200, "ASCII Non-print")
	return err
}
