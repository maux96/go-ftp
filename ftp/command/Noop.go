package commands

type NoopCommand BaseCommand

func (nc NoopCommand) Action(ctx *ConnCtx) error {
	_, err := ctx.conn.Write([]byte("200 OK\t\n"))
	return err
}
