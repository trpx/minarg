package parsers

type Parser interface {
}

type StringParser struct {
}

func (p *StringParser) parse(args []string) (value string, err error) {
	value = args[1]
	return value, nil
}
