package args

import (
	"github.com/trpx/minarg/accumulators"
	"github.com/trpx/minarg/consumers"
	"github.com/trpx/minarg/errors"
	"github.com/trpx/minarg/marshallers"
)

type StringArg struct {
	consumer    consumers.PositionalConsumer
	accumulator accumulators.FixedAccumulator
	marshaller  marshallers.StringMarshaller
	dst         **string
}

func (a *StringArg) Parse(args []string) (remainder []string, err errors.ParseError) {
	if a.dst != nil {
		return args, err
	}
	consumed, remainder := a.consumer.Consume(args)
	if len(consumed) == 0 {
		return args, err
	}
	err = a.accumulator.Accumulate(consumed[0])
	if err != nil {
		return remainder, err
	}
	result, err := a.marshaller.Marshall(consumed[0])
	*a.dst = result
	return args, err
}

func (a *StringArg) Validate() (err errors.ParseError) {
	return err
}
