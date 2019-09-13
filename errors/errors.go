package errors

import "fmt"

type ParseError interface {
	error
	Code() int
}

type parseError struct {
	msg  string
	code int
}

const (
	_ = iota
	TYPE_ERR
	TOO_MANY_VALUES
)

func (p parseError) Error() string {
	return p.msg
}

func (p parseError) Code() int {
	return p.code
}

func NewParseError(code int, format string, a ...interface{}) ParseError {
	return parseError{
		msg:  fmt.Sprintf(format, a...),
		code: code,
	}
}
