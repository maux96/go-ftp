package commands

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type PortCommand BaseCommand

func (cmd PortCommand) Action(ctx *ConnCtx) error {

	rawAddr := strings.Join(cmd.Args, " ")
	addrAsList := strings.Split(rawAddr, ",")
	host := strings.Join(addrAsList[:4], ".")

	// handle posible errors from Atoi
	p1, _ := strconv.Atoi(addrAsList[4])
	p2, _ := strconv.Atoi(addrAsList[5])

	port := p1*256 + p2

	dataConn, err := net.Dial("tcp4", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		// send error response to the client
		return err
	}

	ctx.dataConnection = dataConn
	err = ctx.SendMessage(200, "Entering in Active Mode")
	return err
}
