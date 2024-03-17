package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const (
		address = "127.0.0.1:8000"
		network = "tcp"
	)
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln("can't lister on on given address", address, err)

	}

	fmt.Println("server listening on", listener.Addr())

	for {
		connection, aErr := listener.Accept()
		if aErr != nil {
			log.Println("can't listen to new connection", aErr)

			continue
		}

		// process request
		var data = make([]byte, 1024)
		numberOfReadBytes, rErr := connection.Read(data)
		if rErr != nil {
			log.Println("can't read data from connection", rErr)

			continue
		}

		fmt.Printf("client address: %s , numberOfReadBytes: %d, data: %s\n",
			connection.RemoteAddr(), numberOfReadBytes, data)

		_, wErr := connection.Write([]byte(`your message received`))
		if wErr != nil {
			log.Println("can't write data to connection", wErr)

			continue
		}
	}

}
