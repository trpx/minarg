package consumers

type Consumer interface {
	Consume()
}

type PositionalConsumer struct {
}

func (c *PositionalConsumer) Consume(args []string) (consumed []string, remainder []string) {
	for _, arg := range args {
		firstChar := arg[:1]
		if firstChar != "-" {
			consumed = append(consumed, arg)
		} else {
			break
		}
	}
	return consumed, remainder
}
