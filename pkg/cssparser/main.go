package cssparser

import (
	"errors"
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

type ParseResult struct {
	Output   []byte
	Valid    bool
	Errors   []error
	Warnings []string
}

func wrapErr(message string, err error) error {
	if err == nil {
		return errors.New(message)
	}
	return fmt.Errorf("%s; %w", message, err)
}

func Parse(styles []byte, errIfInvalid bool) (result *ParseResult, err error) {

	buildOpts := api.BuildOptions{
		Stdin: &api.StdinOptions{
			Contents: string(styles),
			Loader:   api.LoaderCSS,
		},
		Format: api.FormatDefault,
	}

	r := api.Build(buildOpts)

	warnings := make([]string, len(r.Warnings))
	for i, w := range r.Warnings {
		warnings[i] = w.Text
		if errIfInvalid {
			err = wrapErr(w.Text, err)
		}
	}

	errs := make([]error, len(r.Errors))
	for i, e := range r.Errors {
		errs[i] = errors.New(e.Text)
		err = wrapErr(e.Text, err)
	}

	result = &ParseResult{
		Output:   r.OutputFiles[0].Contents,
		Valid:    len(errs) == 0 && len(warnings) == 0,
		Errors:   errs,
		Warnings: warnings,
	}

	return
}
