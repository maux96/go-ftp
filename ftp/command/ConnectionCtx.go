package commands

import (
	"net"
)

type ConnCtx struct {
	conn        net.Conn
	currentPath string
}

func NewConnCtx(conn net.Conn) *ConnCtx {
	return &ConnCtx{
		conn:        conn,
		currentPath: "/home", // change by some default
	}
}
