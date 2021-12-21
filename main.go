package main

import (
	"os"

	"github.com/masterpug99/learnblockgo/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.run()
}
