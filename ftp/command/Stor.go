package commands

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type StorCommand BaseCommand

func (cmd StorCommand) Action(ctx *ConnCtx) error {

	path := strings.Join(cmd.Args, " ")
	fileName := filepath.Base(path)
	log.Println(path)
	parentPath := path[:len(path)-len(fileName)]
	log.Println(parentPath)

	realParentPath, exists := ctx.GetRealPath(parentPath)
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

	fd, err := os.OpenFile(filepath.Join(realParentPath, fileName), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fd.Close()

	var buffer [2048]byte
	for {
		n, err_ := ctx.dataConnection.Read(buffer[:])
		if err_ != nil {
			break
		}
		_, err_ = fd.Write(buffer[:n])
		if err_ != nil {
			return err_
		}
	}
	log.Printf("File %s strored!\n", path)
	ctx.dataConnection.Close()

	err = ctx.SendMessage(226, "Closing data connection.")
	return err
}
