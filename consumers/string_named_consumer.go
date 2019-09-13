package consumers

import (
	"github.com/trpx/minarg/constants"
)

type StringNamedConsumer struct {
	argName string
}

func (c *StringNamedConsumer) Consume(args []string) (consumed []string, remainder []string) {
	if len(args) < 2 {
		return consumed, args
	}

	arg := args[0]
	value := args[1]

	if arg == constants.PREFIX+c.argName {
		consumed = append(consumed, value)
		remainder = args[2:]
	} else {
		remainder = args
	}
	return consumed, remainder
}
