package commands

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

var ROOT_PATH string = "/home/maux96/Images/Space"

type ConnCtx struct {
	conn        net.Conn
	currentPath string
	basePath    string
}

func NewConnCtx(conn net.Conn) *ConnCtx {
	return &ConnCtx{
		conn:        conn,
		currentPath: ROOT_PATH,
		basePath:    ROOT_PATH,
	}
}

func (ctx *ConnCtx) SendMessage(code int, message string) error {
	_, err := ctx.conn.Write([]byte(fmt.Sprintf("%d %s%s", code, message, "\t\n")))
	return err
}

func (ctx *ConnCtx) UserPath() string {
	sol, err := filepath.Rel(ctx.basePath, ctx.currentPath)
	if err != nil {
		log.Fatal("Este error no debio suceder,", err.Error())
	}
	return sol
}
func (ctx *ConnCtx) ChangePath(newPath string) error {
	var realNewPath string
	var tempPath string
	if filepath.IsAbs(newPath) {
		tempPath = filepath.Clean(newPath)
	} else {
		tempPath = filepath.Clean(filepath.Join(ctx.UserPath(), newPath))
	}

	if filepath.Base(tempPath) == ".." {
		realNewPath = ctx.basePath
	} else {
		realNewPath = filepath.Join(ctx.basePath, tempPath)
	}

	if ok, err := pathExists(realNewPath); ok {
		ctx.currentPath = realNewPath
		return nil
	} else {
		return err
	}
}
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, err
	} else {
		return false, err
	}
}
