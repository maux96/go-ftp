package commands

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type RnfrCommand BaseCommand
type RntoCommand BaseCommand

func (cmd RnfrCommand) Action(ctx *ConnCtx) error {
	path := strings.Join(cmd.Args, " ")

	realPath, exists := ctx.GetRealPath(path)
	if !exists {
		err := ctx.SendMessage(501, "Invalid path.")
		return err
	}

	ctx.SetRenameFromPath(realPath)

	err := ctx.SendMessage(350, "Requested file action pending further information.")
	return err
}

func (cmd RntoCommand) Action(ctx *ConnCtx) error {
	pathTo := strings.Join(cmd.Args, " ")

	realPathFrom, ok := ctx.GetRenameFromPath()
	if !ok {
		err := ctx.SendMessage(503, "Bad sequence of commands.")
		return err
	}

	name := filepath.Base(pathTo)
	parentPath := pathTo[:len(pathTo)-len(name)]
	_, existsParentPath := ctx.GetRealPath(parentPath)
	if !existsParentPath {
		err := ctx.SendMessage(501, "Parent path not exists!")
		return err
	}

	realPathTo, existsPathTo := ctx.GetRealPath(pathTo)

	if existsPathTo {
		err := ctx.SendMessage(501, "Path exists!")
		return err
	}

	err := os.Rename(realPathFrom, realPathTo)
	if err != nil {
		errSending := ctx.SendMessage(501, "Path exists!")
		return errors.Join(err, errSending)
	}

	err = ctx.SendMessage(250, "Requested file action okay, completed.")
	return err
}
