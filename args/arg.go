package args

import (
	"github.com/trpx/minarg/errors"
)

type ArgParser interface {
	Parse(args []string) (remainder []string, err errors.ParseError)
	Validate() (err errors.ParseError)
}
