package commands

type UserCommand BaseCommand

func (cmd UserCommand) Action(ctx *ConnCtx) error {
	err := ctx.SendMessage(230, "User anon")
	return err
}
