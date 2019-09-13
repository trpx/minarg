package accumulators

import (
	"github.com/trpx/minarg/errors"
)

// Used to accumulate results of consumers
type Accumulator interface {
	Accumulate(arg string) (err errors.ParseError)
	Result() []string
}

// todo: refactor to minCapacity, maxCapacity
type FixedAccumulator struct {
	Capacity int
	result   []string
}

func (p *FixedAccumulator) Accumulate(arg string) (err errors.ParseError) {
	if len(p.result)+1 > p.Capacity {
		err = errors.NewParseError(
			errors.TOO_MANY_VALUES,
			"too many values",
		)
	} else {
		p.result = append(p.result, arg)
	}
	return err
}

// todo: add err return here too to work with the above minCapacity, maxCapacity
func (p *FixedAccumulator) Result() []string {
	return p.result
}
