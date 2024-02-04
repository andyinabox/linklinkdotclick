package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yargevad/filepathx"
)

func main() {
	var err error
	var glob string

	flag.StringVar(&glob, "g", "", "glob to find infiles")
	flag.Parse()

	if glob == "" {
		fmt.Println("usage: ./cmd/clean -g='glob/**/*.ext'")
		os.Exit(0)
	}

	// get file paths using glob
	var files []string
	files, err = filepathx.Glob(glob)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, f := range files {
		err = os.RemoveAll(f)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

}
