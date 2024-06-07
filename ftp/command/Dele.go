package commands

import (
	"log"
	"os"
	"strings"
)

type DeleCommand BaseCommand

func (cmd DeleCommand) Action(ctx *ConnCtx) error {

	path := strings.Join(cmd.Args, " ")

	realPath, exists := ctx.GetRealPath(path)
	if !exists {
		ctx.SendMessage(451, "Requested action aborted: local error in processing.")
		return nil
	}

	err := os.Remove(realPath)
	if err != nil {
		err = ctx.SendMessage(550, "Requested action not taken. File unavailable.")
		return err
	}

	err = ctx.SendMessage(250, "File removed")
	if err != nil {
		return err
	}

	log.Printf("File %s removed!\n", path)
	return nil
}
