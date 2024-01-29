package main

import (
	"os"

	"github.com/evanw/esbuild/pkg/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
