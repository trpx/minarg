package consumers

type StringNamedConsumer struct {
}

func (c *StringNamedConsumer) Consume(args []string) (consumed []string, remainder []string) {
	if len(args) < 2 {
		return consumed, args
	}
	args = args[:2]
	prefix1 := args[0][:2]
	prefix2 := args[1][:1]
	thirdChar := args[0][2:3]

	if prefix1 == "--" && thirdChar != "-" && prefix2 != "-" {
		consumed = args[:2]
		remainder = args[2:]
	} else {
		remainder = args
	}
	return consumed, remainder
}
