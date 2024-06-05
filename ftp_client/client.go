package FtpClient

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func StartClient() {
	var (
		CONNECTION_ADDR = "localhost"
		CONNECTION_PORT = 21
	)

	if len(os.Args) > 1 {
		CONNECTION_ADDR = os.Args[1]
	}
	if len(os.Args) > 2 {
		portToUse, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			CONNECTION_PORT = portToUse
		}
	}

	addr := fmt.Sprintf("%s:%d", CONNECTION_ADDR, CONNECTION_PORT)

	conn, someError := net.Dial("tcp", addr)
	defer func() {
		if someError != nil {
			log.Fatal("ERROR :", someError.Error())
		}
		conn.Close()
	}()

	if someError != nil {
		log.Fatalln("Error connection to", addr, "Exiting the application")
		return
	}

	var data [1024]byte
	_, someError = conn.Read(data[:])
	if someError != nil {
		return
	}

	fmt.Printf("Welcome Message: %s\n", data)

	scanner := bufio.NewScanner(os.Stdin)

	for {

		scanner.Scan()
		input := scanner.Text()
		println(">>>>", input)

		_, someError = conn.Write([]byte(input))
		if someError != nil {
			return
		}

		var response [1024]byte
		_, someError = conn.Read(response[:])
		if someError != nil {
			return
		}
		log.Printf("%s", response)
	}

}
