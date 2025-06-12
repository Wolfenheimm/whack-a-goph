package main

import (
	"os"

	"github.com/raithbheart/whack-a-goph/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
