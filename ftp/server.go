package FtpServer

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type FtpServer struct {
	host string
	port int

	welcomeMessage string
}


func New(host string, port int) *FtpServer{
	server := FtpServer{
		host: host,
		port: port,
		welcomeMessage: "Default Welcome Message :D",
	}

	return &server
}

func (server *FtpServer) Run() (err error) {
	log.Println("Server staring...")

	listening, err:= net.Listen("tcp", server.host+":"+ fmt.Sprint(server.port))
	defer listening.Close()

	if err != nil {
		log.Fatal("Error Starting the server:", err.Error())
	}

	for {
		conn, err := listening.Accept()
		if err != nil {
			return err	
		}

		go server.handleConnection(conn)						
	}
}


func (server *FtpServer) handleConnection(conn net.Conn){
	var posibleErr error
	defer func(){
		if posibleErr != nil{
			log.Println("Error:",posibleErr.Error())
		}
		conn.Close()
	}()
	_,posibleErr=conn.Write([]byte(server.welcomeMessage))
	if posibleErr != nil{
		return
	}


	for  {

		var currentData [1024]byte
		totalDataReaded, posibleErr:=conn.Read(currentData[:]) 
		if posibleErr != nil ||  totalDataReaded == 0 {
			return 
		}

		log.Printf("Reciving command from %s : %s", conn.RemoteAddr(), currentData)	

		currentDataAsString:=strings.Trim(string(currentData[:]),"\x00")
		tokens:=strings.Split(currentDataAsString," ")
		tokens[0] =  strings.ToUpper(tokens[0])
		fmt.Println(tokens)
		
		switch tokens[0]{

		case "NOOP":
			_,posibleErr=conn.Write([]byte("200 OK\t\n"))	
			if posibleErr != nil {
				return	
			}
		default:
			_,posibleErr=conn.Write([]byte("502 Command Not Implemented!\t\n"))		
			if posibleErr != nil{
				return
			}
		}
	}
}




