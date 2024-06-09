package commands

import (
	"os"
	"path/filepath"
	"strings"
)

type MkdCommand BaseCommand

func (cmd MkdCommand) Action(ctx *ConnCtx) error {
	/* TODO verifiy when exist a file named like the directory */

	path := strings.Join(cmd.Args, " ")
	dirName := filepath.Base(path)
	parentPath := strings.Join(cmd.Args, " ")[:len(path)-len(dirName)]

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
