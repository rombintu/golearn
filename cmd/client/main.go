package main

import (
	"log"
	"os"

	"github.com/rombintu/golearn/external"
)

func main() {
	terminal := external.NewTerminal("golearn", "Test usage")
	if err := terminal.Building(); err != nil {
		log.Fatal(err)
	}
	err := terminal.CLI.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
