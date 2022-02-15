package main

import (
	"flag"
	"log"
	"os"

	"github.com/rombintu/golearn/config"
	"github.com/rombintu/golearn/external"
	"github.com/rombintu/golearn/internal/client"
	"github.com/urfave/cli"
)

func buildClientCLI(term *external.Terminal, conf *config.Config) {
	// term.InterStore["host"] = conf.Client.Host
	// term.InterStore["port"] = conf.Client.Port
	flags := []cli.Flag{
		cli.StringFlag{
			Name:     "login",
			Usage:    "Your login",
			Required: true,
		},
		cli.StringFlag{
			Name:     "pass",
			Usage:    "Your password",
			Required: true,
		},
		cli.StringFlag{
			Name:     "role",
			Usage:    "Your role",
			Required: false,
		},
	}
	term.AddCommand(
		cli.Command{
			Name:  "ping",
			Usage: "Ping to server",
			Action: func(c *cli.Context) error {
				ping, err := client.PingServer(
					conf.Client.Host + conf.Client.Port,
				)
				if err != nil {
					return err
				}
				term.Output.Info(ping)
				return nil
			},
		},
	)
	term.AddCommand(
		cli.Command{
			Name:  "auth",
			Usage: "Authentification (get token)",
			Flags: flags,
			Action: func(c *cli.Context) error {
				token, err := client.GetToken(
					conf.Client.Host+conf.Client.Port,
					c.String("login"),
					c.String("pass"),
					c.String("role"),
				)
				if err != nil {
					return err
				}
				term.Output.Info(token)
				return nil
			},
		},
	)
}

func main() {
	configPath := flag.String("config", "./config/client.toml", "Path to config file")
	flag.Parse()

	config := config.GetConfig(*configPath)

	terminal := external.NewTerminal("Golearncli", "Golearn CLI")
	if err := terminal.Building(); err != nil {
		log.Fatal(err)
	}

	buildClientCLI(terminal, config)

	err := terminal.CLI.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
