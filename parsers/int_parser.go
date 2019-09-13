package parsers

import (
	"github.com/trpx/minarg/errors"
	"strconv"
)

type IntParser struct {
}

func (p *IntParser) parse(arg string) (result *int, err errors.ParseError) {
	integer, convertErr := strconv.Atoi(arg)
	if convertErr != nil {
		err = errors.NewParseError(
			errors.TYPE_ERR,
			"not a valid int",
		)
	} else {
		result = &integer
	}
	return result, err
}
