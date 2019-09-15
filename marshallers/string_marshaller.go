package marshallers

import (
	"github.com/trpx/minarg/errors"
)

type Marshaller interface {
	Marshall(value string) (err errors.ParseError)
}

type StringMarshaller struct {
}

func (p *StringMarshaller) Marshall(arg string) (result *string, err errors.ParseError) {
	result = &arg
	err = nil
	return result, err
}
