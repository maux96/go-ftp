package commands

import (
	"os"
	"strings"
)

type MkdCommand BaseCommand

func (cmd MkdCommand) Action(ctx *ConnCtx) error {
	/* TODO verifiy when exist a file named like the directory */
	parentPath := strings.Join(cmd.Args[:len(cmd.Args)-1], " ")
	path := strings.Join(cmd.Args, " ")

	_, exists := ctx.GetRealPath(parentPath)
	if !exists {
		err := ctx.SendMessage(501, "Invalid Directory.")
		return err
	}

	realNewDirPath, exists := ctx.GetRealPath(path)
	if exists {
		err := ctx.SendMessage(501, "The current directory exists")
		return err
	}
	err := os.Mkdir(realNewDirPath, os.ModePerm)
	if err != nil {
		ctx.SendMessage(257, "Directory Created.")
		return err
	}

	err = ctx.SendMessage(257, "Directory Created.")
	return err
}
