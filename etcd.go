package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.etcd.io/etcd/client"
)

var initEtcdHandler sync.Once
var etcdGlobalHandler client.KeysAPI

func getEtcdHandler() client.KeysAPI {
	initEtcdHandler.Do(func() {
		cfg := client.Config{
			//Endpoints: []string{"http://docker.for.mac.host.internal:2379"},
			Endpoints: []string{"http://127.0.0.1:2379", "http://127.0.0.1:2380"},
			Transport: client.DefaultTransport,
			// set timeout per request to fail fast when the target endpoint is unavailable
			HeaderTimeoutPerRequest: 10 * time.Second,
			/* use default SelectionMode as EndpointSelectionRandom*/
		}
		c, err := client.New(cfg)
		if err != nil {
			log.Printf("create client to %v error \n", cfg)
			log.Fatal(err)
		}

		etcdGlobalHandler = client.NewKeysAPI(c)
	})

	return etcdGlobalHandler
}

func listDebug(kapi client.KeysAPI) error {
	resp, err := kapi.Get(context.Background(), "/", nil)
	if err != nil {
		log.Printf("Tring GetAllModuletype error \n")
		return err
	}
	for _, node := range resp.Node.Nodes {
		if node.Dir {
			fmt.Printf("found dirs:%s\n", node.Key)
		} else {
			fmt.Printf("%s:%s\n", node.Key, node.Value)

		}
	}
	return nil
}
