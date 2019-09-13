package consumers

// The purpose of consumer is to pop all arguments it can
// from the left hand side of os.Args and return them and remainder
// for further usage by other parts of the library
type Consumer interface {
	Consume()
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
