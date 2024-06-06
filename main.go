package main

import (
	ftp "ftp/ftp"
	client "ftp/ftp_client"
	"os"
)

const (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = 2021
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
