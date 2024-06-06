package commands

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type ListCommand BaseCommand

func (cmd ListCommand) Action(ctx *ConnCtx) error {
	path := strings.Join(cmd.Args, " ")
	realPath, exists := ctx.GetRealPath(path)
	if !exists {
		ctx.SendMessage(451, "Requested action aborted: local error in processing.")
		return nil
	}

	ctx.SendMessage(125, "Data connection already open; transfer starting.")

	files, err := os.ReadDir(realPath)
	if err != nil {
		return err
	}

	toSend := fmt.Sprintf("total %d\n", len(files))
	for _, val := range files {
		info, fileErr := val.Info()
		if fileErr == nil {
			toSend += fmt.Sprintf("%s 1 anon anon %d %s\n", info.Mode().String(), info.Size(), info.Name())
		}
	}

	log.Println(toSend)
	_, err = ctx.dataConnection.Write([]byte(toSend))
	if err != nil {
		return err
	}
	ctx.dataConnection.Close()

	err = ctx.SendMessage(226, "Closing data connection.")
	return err
}
