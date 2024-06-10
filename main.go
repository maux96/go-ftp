package main

import (
	"flag"
	ftp "ftp/ftp"
	"log"
)

var (
	SERVER_HOST = "0.0.0.0"
	SERVER_PORT = 2021
)

func main() {
	flag.StringVar(&SERVER_HOST, "host", SERVER_HOST, "FTP server host (default: 0.0.0.0).")
	flag.IntVar(&SERVER_PORT, "port", SERVER_PORT, "FTP server main port (default: 2021).")
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalln("First argument must be the base path!")
	}
	ROOT_PATH := flag.Arg(0)

	server := ftp.New(SERVER_HOST, SERVER_PORT, ROOT_PATH)
	server.Run()
}
