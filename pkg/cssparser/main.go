package cssparser

import (
	"errors"
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

type ParseOptions struct {
}

func Parse(styles []byte, opts *ParseOptions) (output []byte, valid bool, err error) {

	buildOpts := api.BuildOptions{
		Stdin: &api.StdinOptions{
			Contents: string(styles),
			Loader:   api.LoaderCSS,
		},
		Format: api.FormatDefault,
	}

	result := api.Build(buildOpts)

	if len(result.Errors) > 0 {
		err = fmt.Errorf("errors: %v", result.Errors)
		valid = false
		return
	}

	if len(result.OutputFiles) == 0 {
		err = errors.New("no output files")
		valid = false
		return
	}

	valid = len(result.Warnings) < 1

	output = result.OutputFiles[0].Contents
	return
}
