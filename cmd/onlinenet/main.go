package main

import (
	"fmt"
	"os"
	"path"
	"text/tabwriter"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/moul/onlinenet/pkg/api"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul"
	app.Version = "dev"
	app.Usage = "Client for api.online.net"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Value:  "test",
			Usage:  "API token",
			EnvVar: "ONLINENET_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "servers",
			Action: actionServers,
		},
	}

	app.Run(os.Args)
}

func actionServers(c *cli.Context) {
	client := api.NewClientWithToken(c.GlobalString("token"))

	serverPaths, err := client.ListServers()
	if err != nil {
		logrus.Fatalf("Cannot list servers: %v", err)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 1, 3, ' ', 0)

	fmt.Fprintln(w, "ID\tName\tOs\tOffer")

	for _, serverPath := range *serverPaths {
		server, err := serverPath.Get(client)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", server.Identifier, server.Hostname, server.Os.Name, server.Offer)
	}

	w.Flush()
}
