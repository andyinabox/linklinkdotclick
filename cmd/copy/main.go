package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/yargevad/filepathx"
)

func copy(filepath string, outdir string) (err error) {
	outpath := path.Join(outdir, path.Base(filepath))
	fmt.Printf("copying from %s to %s\n", filepath, outpath)

	var content []byte
	content, err = ioutil.ReadFile(filepath)
	if err != nil {
		return
	}

	err = os.MkdirAll(outdir, os.ModePerm)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(outpath, content, os.ModePerm)
	return
}

func main() {
	var err error
	var out string
	var glob string

	flag.StringVar(&out, "o", "", "outdir")
	flag.StringVar(&glob, "g", "", "glob to find infiles")
	flag.Parse()

	if glob == "" || out == "" {
		fmt.Println("usage: ./cmd/copy -g='glob/**/*.ext' out=path/to/outdir")
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
		err = copy(f, out)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

}
