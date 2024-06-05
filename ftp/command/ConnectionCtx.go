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
	conn           net.Conn
	currentPath    string
	basePath       string
	dataConnection net.Conn
}

func NewConnCtx(conn net.Conn) *ConnCtx {
	return &ConnCtx{
		conn:           conn,
		currentPath:    ROOT_PATH,
		basePath:       ROOT_PATH,
		dataConnection: nil,
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

/* the real path and if exists */
func (ctx *ConnCtx) GetRealPath(pathFromUser string) (string, bool) {
	var realNewPath string
	var tempPath string
	if filepath.IsAbs(pathFromUser) {
		tempPath = filepath.Clean(pathFromUser)
	} else {
		tempPath = filepath.Clean(filepath.Join(ctx.UserPath(), pathFromUser))
	}

	if filepath.Base(tempPath) == ".." {
		realNewPath = ctx.basePath
	} else {
		realNewPath = filepath.Join(ctx.basePath, tempPath)
	}

	exists, _ := pathExists(realNewPath)

	return realNewPath, exists
}

func (ctx *ConnCtx) ChangePath(newPath string) error {
	realNewPath, _ := ctx.GetRealPath(newPath)

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
