package consumers

type Consumer interface {
	Consume()
}

type PositionalConsumer struct {
}

func (c *PositionalConsumer) Consume(args []string) (consumed []string, remainder []string) {
	for idx, arg := range args {
		firstChar := arg[:1]
		if firstChar != "-" {
			consumed = append(consumed, arg)
		} else {
			remainder = args[idx:]
			break
		}
	}
	return consumed, remainder
}
