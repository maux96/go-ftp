package commands

import (
	"fmt"
	"io/fs"
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
	_renameFrom    *string
}

func NewConnCtx(conn net.Conn) *ConnCtx {
	return &ConnCtx{
		conn:           conn,
		currentPath:    ROOT_PATH,
		basePath:       ROOT_PATH,
		dataConnection: nil,
		_renameFrom:    nil,
	}
}

func (ctx *ConnCtx) GetRenameFromPath() (path string, ok bool) {
	if ctx._renameFrom != nil {
		return *ctx._renameFrom, true
	}
	return "", false
}
func (ctx *ConnCtx) SetRenameFromPath(path string) {
	ctx._renameFrom = &path
	log.Println(*ctx._renameFrom)
}

func (ctx *ConnCtx) SendMessage(code int, message string) error {
	_, err := ctx.conn.Write([]byte(fmt.Sprintf("%d %s%s", code, message, " \t\n")))
	return err
}

func (ctx *ConnCtx) UserPath() string {
	sol, err := filepath.Rel(ctx.basePath, ctx.currentPath)
	if err != nil {
		log.Fatal("Este error no debio suceder,", err.Error())
	}
	return filepath.Join("/", sol)

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

	exists := PathExists(realNewPath)

	return realNewPath, exists
}

func (ctx *ConnCtx) ChangePath(newPath string) error {
	realNewPath, _ := ctx.GetRealPath(newPath)

	if ok := PathExists(realNewPath); ok {
		ctx.currentPath = realNewPath
		return nil
	} else {
		return fs.ErrNotExist
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
