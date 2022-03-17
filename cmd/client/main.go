package main

import (
	"log"
	"os"

	"github.com/rombintu/golearn/config"
	"github.com/rombintu/golearn/external"
	"github.com/rombintu/golearn/external/client"
	"github.com/urfave/cli"
)

func buildClientCLI(term *external.Terminal) {
	flagsAuth := []cli.Flag{
		cli.StringFlag{
			Name:     "role",
			Usage:    "Your role",
			Required: false,
		},
	}
	flagsDeclaration := []cli.Flag{
		cli.StringFlag{
			Name:     "action",
			Usage:    "choose action (create, delete, list)",
			Required: true,
		},
		cli.StringFlag{
			Name:  "title",
			Usage: "choose title declaration",
		},
		cli.StringFlag{
			Name:  "uid",
			Usage: "user id",
		},
	}
	term.AddCommand(
		cli.Command{
			Name:  "ping",
			Usage: "Ping to server",
			Action: func(c *cli.Context) error {
				ping, err := term.Client.PingServer()
				if err != nil {
					return err
				}
				term.Logger.Info(ping)
				return nil
			},
		},
	)
	term.AddCommand(
		cli.Command{
			Name:  "auth",
			Usage: "Authentification (get token)",
			Flags: flagsAuth,
			Action: func(c *cli.Context) error {
				token, err := term.Client.GetToken(
					term.Client.Config.Private.Login,
					term.Client.Config.Private.Password,
					c.String("role"),
				)
				if err != nil {
					return err
				}
				if token != "" {
					term.Logger.Infof("Your token: %s", token)
				} else {
					term.Logger.Error("User not found")
				}

				return nil
			},
		},
	)
	term.AddCommand(
		cli.Command{
			Name:  "declaration",
			Usage: "Declarations manager (create, delete, list)",
			Flags: flagsDeclaration,
			Action: func(c *cli.Context) error {
				action := c.String("action")
				title := c.String("title")
				userID := c.String("uid")
				term.Client.DeclarationAction(action, title, userID)
				return nil
			},
		},
	)
}

func main() {
	// configPath := flag.String("config", "./config/client.toml", "Path to config file")
	// help := flag.Bool("help", false, "Print defaults")
	// flag.Parse()

	// if *help {
	// 	flag.PrintDefaults()
	// 	fmt.Println("  Commands: \n\t- auth\n\t- ping")
	// 	os.Exit(0)
	// }

	// config := config.GetClientConfig(*configPath)
	config := config.GetClientConfig("./config/client.toml")
	client := client.NewClient(config)
	terminal := external.NewTerminal("Golearncli", "Golearn CLI", client)
	if err := terminal.Building(); err != nil {
		log.Fatal(err)
	}

	buildClientCLI(terminal)

	err := terminal.CLI.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
