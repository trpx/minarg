package splitter

import (
	"github.com/trpx/minarg/constants"
	"github.com/trpx/minarg/utils"
	"unicode/utf8"
)

type Splitter struct {
	stringArgNames  []string
	boolArgNames    []string
	subCommandNames []string
}

func (s *Splitter) Split(args []string) (parts Parts, remainder []string) {

	parts.NamedStrings = make(map[string][]string)
	parts.NamedBool = make(map[string][]bool)

	positionalParsed := false

	for len(args) > 0 {
		initialArgsLen := len(args)

		arg := args[0]

		// ** Named args **

		// Commands
		if s.isCommandArg(arg) {
			parts.Subcommand = args
			return parts, remainder
		}
		// Positional
		if s.isPositionalArg(arg) {
			parts.Positional = append(parts.Positional, arg)
			args = args[1:]
			// Long args
		} else if s.isLongArgFlag(arg) {
			positionalParsed = true
			argName := arg[2:]
			if s.isBoolArg(argName) {
				parts.addNamedBool(argName, true)
				args = args[1:]
			} else {
				if len(args) < 2 {
					continue
				}
				argValue := args[1]
				parts.addNamedString(argName, argValue)
				args = args[2:]
			}
			// Short args
		} else if s.isShortArgFlag(arg, args) {
			positionalParsed = true
			runes := utils.Runes(arg)[1:]
			length := len(runes)
			for idx, character := range runes {
				isLast := idx == length-1
				if isLast && !s.isBoolArg(character) {
					parts.addNamedString(character, args[1])
					args = args[2:]
				} else {
					parts.addNamedBool(character, true)
					if isLast {
						args = args[1:]
					}
				}
			}
		}
		if len(args) == initialArgsLen {
			remainder = args
			return parts, remainder
		}
	}

	return parts, remainder
}

func (s *Splitter) isCommandArg(arg string) bool {
	if !s.isPositionalArg(arg) {
		return false
	}
	for _, i := range s.subCommandNames {
		if i == arg {
			return true
		}
	}
	return false
}

func (s *Splitter) isPositionalArg(arg string) bool {
	return arg[:1] != constants.SHORT_PREFIX && arg[:2] != constants.LONG_PREFIX
}

func (s *Splitter) isShortArgFlag(arg string, args []string) bool {
	if utf8.RuneCountInString(arg) < 2 {
		return false
	}

	if arg[:1] != constants.SHORT_PREFIX {
		return false
	}

	hasNext := len(args) > 0

	runes := utils.Runes(arg[1:])
	length := len(runes)
	for idx, character := range runes {
		isLast := idx == length-1
		ok := false
		for _, shortArgName := range s.shortArgNames() {
			if character == shortArgName {
				if !isLast || !hasNext {
					if !s.isBoolArg(character) {
						// not bool arg has no value
						return false
					}
				}
				ok = true
				break
			}
		}
		if !ok {
			// there is no such short arg name defined
			return false
		}
	}
	return true
}

func (s *Splitter) isBoolArg(arg string) bool {
	for _, i := range s.boolArgNames {
		if arg == i {
			return true
		}
	}
	return false
}

func (s *Splitter) isLongArgFlag(arg string) bool {
	for _, i := range s.longArgNames() {
		if constants.LONG_PREFIX+i == arg {
			return true
		}
	}
	return false
}

func (s *Splitter) shortArgNames() (names []string) {
	for _, i := range s.stringArgNames {
		if len(i) == 1 {
			names = append(names, i)
		}
	}
	for _, i := range s.boolArgNames {
		if len(i) == 1 {
			names = append(names, i)
		}
	}
	return names
}

func (s *Splitter) longArgNames() (names []string) {
	for _, i := range s.stringArgNames {
		if len(i) > 1 {
			names = append(names, i)
		}
	}
	for _, i := range s.boolArgNames {
		if len(i) > 1 {
			names = append(names, i)
		}
	}
	return names
}
