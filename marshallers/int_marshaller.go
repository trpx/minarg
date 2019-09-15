package marshallers

import (
	"github.com/trpx/minarg/errors"
	"strconv"
)

type IntMarshaller struct {
}

func (p *IntMarshaller) Marshall(arg string) (result *int, err errors.ParseError) {
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
