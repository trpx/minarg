package consumers

// The purpose of consumer is to pop 1 argument if possible
// from the left hand side of os.Args and return it and remainder
// for further usage by other parts of the library
// consumed is a slice because we need to be able to distinguish
// the cases when nothing is consumed from something consumed
// and that would not be possible if the return would be a string
// max return len of the consumed slice == 1, min == 0
type Consumer interface {
	Consume() (consumed []string, remainder []string)
}

type PositionalConsumer struct {
}

func (c *PositionalConsumer) Consume(args []string) (consumed []string, remainder []string) {
	if len(args) < 1 {
		return consumed, args
	}
	arg := args[0]
	firstChar := arg[:1]
	if firstChar != "-" {
		consumed = append(consumed, arg)
		remainder = args[1:]
	} else {
		remainder = args
	}
	return consumed, remainder
}
