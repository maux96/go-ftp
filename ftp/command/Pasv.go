package commands

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
)

type PasvCommand BaseCommand

func (cmd PasvCommand) Action(ctx *ConnCtx) error {

	serverHost := strings.Split(ctx.conn.LocalAddr().String(), ":")[0]
	serverHost = strings.Join(strings.Split(serverHost, "."), ",")

	listener, port, err := initPasiveDataSocket()
	if err != nil {
		return err
	}
	pasivePortAddr := fmt.Sprintf("%s,%d,%d", serverHost, port/256, port%256)

	err = ctx.SendMessage(227, fmt.Sprintf("Entering Passive Mode (%s)", pasivePortAddr))
	if err != nil {
		return err
	}

	conn, err := listener.Accept()
	if err != nil {
		return err
	}

	ctx.dataConnection = conn
	return err
}

func initPasiveDataSocket() (listener net.Listener, port int, err error) {
	port = 1500 + rand.Intn(7500)
	listener, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	return listener, port, err
}
