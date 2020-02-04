package main

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func show_rec(client net.Conn) {
	var message []byte
	message = make([]byte, 1024)
	log.Println("in show_rec")
	log.Println(client.RemoteAddr())

	ra := client.RemoteAddr()
	vec := strings.Split(ra.String(), ":")
	log.Println(vec[0])

	len, _ := client.Read(message)

	if len > 0 {

		log.Println(message[0:len])
	}

}

func recvMessage(client net.Conn) error {
	var first_read bool
	first_read = true
	for {
		if true == first_read {
			first_read = false
			go show_rec(client)
		}

		client.Write([]byte("Big Dick of Go server , Ass Hole ! \n"))
		time.Sleep(time.Duration(2) * time.Second)
	}

	return nil
}

func main() {
	server, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatal("start server failed!\n")
		os.Exit(1)
	}
	defer server.Close()

	log.Println("server is running...")
	for {
		client, err := server.Accept()
		if err != nil {
			log.Fatal("Accept error\n")
			continue
		}

		log.Println("the client is connectted...")
		go recvMessage(client)
	}
}
