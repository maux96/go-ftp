package FtpClient

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	CONNECTION_ADDR = "localhost"
	CONNECTION_PORT = 7000
)

func StartClient() {
	var name string

	//fmt.Print("Your Name:")
	//_,someError:=fmt.Scanln(&name)
	var someError error
	if someError != nil {
		log.Fatalln("Error reading the name!")
	}

	addr := fmt.Sprintf("%s:%d", CONNECTION_ADDR, CONNECTION_PORT)
	log.Println("Connecting", name, "to", addr)

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
