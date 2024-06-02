package main

import (
	ftp "ftp/ftp"
	client "ftp/ftp_client"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = 7000
)

func main() {
	isClient := true
	if len(os.Args) > 1 && os.Args[1] == "server" {
		isClient = false
	}

	if isClient {
		client.StartClient()
	} else {
		server := ftp.New(SERVER_HOST, SERVER_PORT)
		server.Run()
	}
}
