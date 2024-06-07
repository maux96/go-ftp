package commands

import (
	"log"
	"os"
	"strings"
)

type RetrCommand BaseCommand

func (cmd RetrCommand) Action(ctx *ConnCtx) error {

	path := strings.Join(cmd.Args, " ")

	realPath, exists := ctx.GetRealPath(path)
	if !exists {
		ctx.SendMessage(451, "Requested action aborted: local error in processing.")
		return nil
	}

	if ctx.dataConnection == nil {
		ctx.SendMessage(426, "Connection closed; transfer aborted.")
		return nil
	}
	defer ctx.dataConnection.Close()

	ctx.SendMessage(125, "Data connection already open; transfer starting.")

	fd, err := os.Open(realPath)
	if err != nil {
		return err
	}
	defer fd.Close()

	var buffer [2048]byte
	for {
		n, err_ := fd.Read(buffer[:])
		if err_ != nil {
			break
		}
		_, err_ = ctx.dataConnection.Write(buffer[:n])
		if err_ != nil {
			return err_
		}
	}
	log.Printf("File %s sended to user!\n", path)
	ctx.dataConnection.Close()

	err = ctx.SendMessage(226, "Closing data connection.")
	return err
}
