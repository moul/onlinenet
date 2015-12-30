package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
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
			Value:  "",
			Usage:  "API token",
			EnvVar: "ONLINENET_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "servers",
			Action: actionServers,
		},
		{
			Name:   "server-show",
			Action: actionServerShow,
		},
		{
			Name:   "server-reboot",
			Action: actionServerReboot,
		},
		{
			Name:   "abuses",
			Action: actionAbuses,
		},
		{
			Name:   "user",
			Action: actionUser,
		},
		{
			Name:   "ddos",
			Action: actionDdos,
		},
	}

	app.Run(os.Args)
}

func actionAbuses(c *cli.Context) {
	client := api.NewClientWithToken(c.GlobalString("token"))

	abuses, err := client.ListAbuses()
	if err != nil {
		logrus.Fatalf("Cannot list abuses: %v", err)
	}

	for _, abuse := range *abuses {
		fmt.Println(abuse)
	}
}

func actionDdos(c *cli.Context) {
	client := api.NewClientWithToken(c.GlobalString("token"))

	ddosList, err := client.ListDdos()
	if err != nil {
		logrus.Fatalf("Cannot list ddos: %v", err)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 1, 3, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "ID\tTarget\tStart\tEnd\tMitigation\tType\tMax PPS\tMax BPS\tTimeline")

	for _, ddos := range *ddosList {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%d\t%d\t%d items\n", ddos.Identifier, ddos.Target, ddos.Start, ddos.End, ddos.Mitigation, ddos.Type, ddos.MaxPPS, ddos.MaxBPS, len(ddos.Timeline))
	}
}

func actionUser(c *cli.Context) {
	client := api.NewClientWithToken(c.GlobalString("token"))

	user, err := client.GetCurrentUser()
	if err != nil {
		logrus.Fatalf("Cannot get current user: %v", err)
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 10, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tLogin\tName\tCompany")
	fmt.Fprintf(w, "%d\t%s\t%s %s\t%s\n", user.Identifier, user.Login, user.FirstName, user.LastName, user.Company)
	w.Flush()
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

func actionServerShow(c *cli.Context) {
	if len(c.Args()) < 1 {
		logrus.Fatalf("You must specify a server")
	}
	serverID, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		logrus.Fatalf("ServerID %q is not a valid number: %v", c.Args()[0], err)
	}

	client := api.NewClientWithToken(c.GlobalString("token"))

	server, err := client.GetServer(serverID)
	if err != nil {
		logrus.Fatalf("Failed to get server %d: %v", serverID, err)
	}

	body, err := json.MarshalIndent(server, "", "  ")
	if err != nil {
		logrus.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(body))
}

func actionServerReboot(c *cli.Context) {
	if len(c.Args()) < 1 {
		logrus.Fatalf("You must specify a server")
	}
	serverID, err := strconv.Atoi(c.Args()[0])
	if err != nil {
		logrus.Fatalf("ServerID %q is not a valid number: %v", c.Args()[0], err)
	}

	client := api.NewClientWithToken(c.GlobalString("token"))

	ret, err := client.RebootServer(serverID, "no reason", "")
	if err != nil {
		logrus.Fatalf("Failed to reboot server %d: %v", serverID, err)
	}

	if *ret == true {
		logrus.Infof("Server %d rebooted", serverID)
	} else {
		logrus.Fatalf("Server %d failed to reboot (no details)", serverID)
	}
}
