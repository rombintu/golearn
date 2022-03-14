package external

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rombintu/golearn/external/client"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Terminal struct {
	CLI        *cli.App
	Logger     *log.Logger
	InterStore map[string]string
	Client     *client.Client
}

func NewTerminal(progName, progUsage string, config *client.Client) *Terminal {
	termCLI := cli.NewApp()
	termCLI.Name = progName
	termCLI.Usage = progUsage

	logger := log.New()
	logger.SetLevel(log.InfoLevel)

	return &Terminal{
		CLI:        termCLI,
		Logger:     logger,
		InterStore: make(map[string]string),
		Client:     config,
	}
}

func (t *Terminal) getInterStore() error {
	file, err := os.Open("./version.json")
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	var version map[string]string

	if err := json.Unmarshal(data, &version); err != nil {
		return err
	}
	t.InterStore = version
	return nil
}

func (t *Terminal) AddFlag(flag cli.StringFlag) {
	t.CLI.Flags = append(t.CLI.Flags, flag)
}

func (t *Terminal) AddCommand(command cli.Command) {
	t.CLI.Commands = append(t.CLI.Commands, command)
}

func (t *Terminal) Building() error {
	if err := t.getInterStore(); err != nil {
		return err
	}
	t.AddFlag(cli.StringFlag{
		Name:  "debug",
		Value: "on",
	})
	t.AddCommand(cli.Command{
		Name:    "version",
		Usage:   "Show version",
		Aliases: []string{"V"},
		Flags:   t.CLI.Flags,
		Action: func(c *cli.Context) error {
			fmt.Printf("Golearn-client: v%s", t.InterStore["version"])
			fmt.Printf("\nAuthor: %s", t.InterStore["author"])
			return nil
		},
	})
	return nil
}
