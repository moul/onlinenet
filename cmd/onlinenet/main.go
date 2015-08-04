package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/moul/onlinenet/pkg/api"
)

func main() {
	client := api.NewClient()

	serverPaths, err := client.ListServers()
	if err != nil {
		logrus.Fatalf("Cannot list servers: %v", err)
	}

	for _, serverPath := range *serverPaths {
		fmt.Println(serverPath)
		server, err := serverPath.Get(client)
		if err != nil {
			panic(err)
		}
		fmt.Println(server.Hostname)
	}
}
