package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"go.etcd.io/etcd/client"
)

func initEntopy(nc net.Conn, kapi client.KeysAPI) string {
	var message []byte
	message = make([]byte, 1024)
	log.Println(nc.RemoteAddr())

	ra := nc.RemoteAddr()
	vec := strings.Split(ra.String(), ":")
	newkey := "/" + vec[0]
	fmt.Printf("newkey:%s\n", newkey)

	_, err := kapi.Get(context.Background(), newkey, &client.GetOptions{Recursive: true})
	if err != nil {
		log.Printf("Trying Get from  %v error , set it as initial\n", newkey)
		_, err := kapi.Set(context.Background(), newkey, "10000", nil)
		if err != nil {
			log.Printf("kapi.Set %v error :%v\n", newkey, err)
		}
	}

	go func(nc net.Conn) {
		len, _ := nc.Read(message)
		if len > 0 {
			log.Println(message[0:len])
		}
	}(nc)
	return newkey
}

func recvMessage(nc net.Conn, kapi client.KeysAPI) error {
	firstRead := true
	var privatePath string
	for {

		if firstRead {
			firstRead = false
			/*go initEntopy(client, kapi)*/
			privatePath = initEntopy(nc, kapi)
		}

		resp, err := kapi.Get(context.Background(), privatePath, &client.GetOptions{Recursive: true})
		if err != nil {
			log.Printf("Trying privatePath error \n")
			break
		}
		var entropy string

		if resp.Node.Dir {
			fmt.Printf("found dirs:%s\n", resp.Node.Key)
		} else {
			fmt.Printf("%s:%s\n", resp.Node.Key, resp.Node.Value)
			entropy = resp.Node.Value
		}

		message := "Your Current entropy is " + entropy + "\n"
		_, err = nc.Write([]byte(message))
		if err != nil {
			fmt.Printf("Connection error and quit\n")
			break
		}
		time.Sleep(time.Duration(2) * time.Second)

		/* munus and set */
		iEntropy, err := strconv.Atoi(entropy)
		if 0 == iEntropy {
			break
		}
		iEntropy--
		sEntropy := strconv.Itoa(iEntropy)

		kapi.Set(context.Background(), privatePath, sEntropy, nil)
	}

	return nil
}

func main() {

	kapi := getEtcdHandler()
	listDebug(kapi)

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
		go recvMessage(client, kapi)
	}

}
