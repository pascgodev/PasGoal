package main

import (
	"github.com/pasgo/pasgo/core/p2p"
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Name = "Pasgoal"
	app.Usage = "More than Golang implement of PascalCoin"
	app.Version = "0.0.1"
	//app.HideHelp = true

	app.Action = func(c *cli.Context) error {
		// main
		port := c.Int("port")
		p := p2p.InitPeer(port)
		p.Start()

		return nil
	}

	flags := []cli.Flag{
		cli.BoolFlag{
			Name:  "daemon, d",
			Usage: "Start PasGoal as a Daemon (background)",
		},
		cli.BoolTFlag{
			Name:  "save-log",
			Usage: "Whether save the log into file",
		},
		cli.IntFlag{
			Name:  "port",
			Usage: "Port for P2P connection",
			Value: 5005,
		},
		cli.IntFlag{
			Name:  "rpc-port",
			Usage: "Port for RPC",
			Value: 5006,
		},
	}

	app.Flags = flags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
