package splitter

type Parts struct {
	Positional   []string
	NamedStrings map[string][]string // flag-name: [value1, value2, ...]
	NamedBool    map[string][]bool
	Subcommand   []string // args starting with subcommand name
}

type namedArg struct {
	name   string
	values []string
}

func (p *Parts) addNamedString(name string, value string) {
	_, ok := p.NamedStrings[name]
	if ok {
		p.NamedStrings[name] = append(p.NamedStrings[name], value)
	} else {
		p.NamedStrings[name] = []string{value}
	}
}

func (p *Parts) getNamedString(name string) (values []string, ok bool) {
	v, ok := p.NamedStrings[name]
	return v, ok
}

func (p *Parts) addNamedBool(name string, value bool) {
	_, ok := p.NamedBool[name]
	if ok {
		p.NamedBool[name] = append(p.NamedBool[name], value)
	} else {
		p.NamedBool[name] = []bool{value}
	}
}
