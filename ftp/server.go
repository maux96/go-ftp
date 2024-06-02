package FtpServer

import (
	"fmt"
	"log"
	"net"
	"strings"

	commands "ftp/ftp/command"
)

type FtpServer struct {
	host string
	port int

	welcomeMessage string
}

func New(host string, port int) *FtpServer {
	server := FtpServer{
		host:           host,
		port:           port,
		welcomeMessage: "Default Welcome Message :D",
	}

	return &server
}

func (server *FtpServer) Run() (err error) {
	log.Println("Server staring...")

	listening, err := net.Listen("tcp", server.host+":"+fmt.Sprint(server.port))
	if err != nil {
		log.Fatal("Error Starting the server:", err.Error())
		return err
	}
	defer listening.Close()

	for {
		conn, err := listening.Accept()
		if err != nil {
			return err
		}

		go server.handleConnection(conn)
	}
}

func (server *FtpServer) handleConnection(conn net.Conn) {
	var posibleErr error
	defer func() {
		if posibleErr != nil {
			log.Println("Error:", posibleErr.Error())
		}
		conn.Close()
	}()
	_, posibleErr = conn.Write([]byte(server.welcomeMessage))
	if posibleErr != nil {
		log.Println("Error writing to " + conn.RemoteAddr().String() + " connection.\n" + posibleErr.Error())
		return
	}

	ctx := commands.NewConnCtx(conn)

	for {
		var currentData [1024]byte
		totalDataReaded, posibleErr := conn.Read(currentData[:])
		if posibleErr != nil || totalDataReaded == 0 {
			log.Println("Error reading from " + conn.RemoteAddr().String() + " connection.\n" + posibleErr.Error())
			return
		}

		log.Printf("Reciving command from %s : %s", conn.RemoteAddr(), currentData)

		var baseCommand commands.BaseCommand = commands.RawCommand(strings.Trim(string(currentData[:]), "\x00")).GetCommand()

		comm, err := ResolveCommand(baseCommand)
		if err != nil {
			log.Println(err.Error())
			conn.Write([]byte("502 Command Not Implemented!\t\n"))
		} else {
			comm.Action(ctx)
		}
	}
}
