package parsers

import (
	"github.com/trpx/minarg/errors"
)

type Parser interface {
	parse(value string) (err error)
}

type StringParser struct {
}

func (p *StringParser) parse(arg string) (result *string, err errors.ParseError) {
	result = &arg
	err = nil
	return result, err
}
