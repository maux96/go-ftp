package FtpServer

import (
	"fmt"
	"log"
	"net"
	"strings"

	commands "ftp/ftp/command"
)

type FtpServer struct {
	host     string
	port     int
	basePath string

	welcomeMessage string
}

func New(host string, port int, basePath string) *FtpServer {
	server := FtpServer{
		host:           host,
		port:           port,
		basePath:       basePath,
		welcomeMessage: "Default Welcome Message :D",
	}

	return &server
}

func (server *FtpServer) Run() (err error) {
	log.Println("Server staring...")

	listening, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", server.host, server.port))
	if err != nil {
		log.Fatal("Error Starting the server:", err.Error())
		return err
	}
	defer listening.Close()

	log.Printf("Server available in %s:%d\n", server.host, server.port)
	log.Printf("Base Path: %s\n", server.basePath)

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

	ctx := commands.NewConnCtx(conn, server.basePath)

	posibleErr = ctx.SendMessage(220, server.welcomeMessage)
	if posibleErr != nil {
		log.Println("Error writing to " + conn.RemoteAddr().String() + " connection.\n" + posibleErr.Error())
		return
	}

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
			err = comm.Action(ctx)
			if err != nil {
				log.Println("ERROR: " + err.Error())
			}
		}
	}
}
