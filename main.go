package main

import (
	ftp "ftp/ftp"
)

const (
	SERVER_HOST="localhost"
	SERVER_PORT=7000
)

func main(){
	server:=ftp.New(SERVER_HOST, SERVER_PORT)
	server.Run()
}
