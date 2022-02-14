package test

import (
	"os"
	"testing"

	"github.com/rombintu/golearn/external"
)

func TestCli(t *testing.T) {
	os.Args = []string{"version"}
	terminal := external.NewTerminal("Test terminal", "Test usage")
	err := terminal.CLI.Run(os.Args)
	if err != nil {
		t.Fatal(err)
	}
}
