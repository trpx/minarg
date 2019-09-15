package parsers

import (
	"github.com/trpx/minarg/args"
	"github.com/trpx/minarg/errors"
)

type ParserController struct {
	subparsers []ParserController
	positional []args.ArgParser
	named      []args.ArgParser
}

func (p *ParserController) Parse(args []string) (remainder []string, err errors.ParseError) {

	args, err = p.parsePositional(args)
	if err != nil {
		return args, err
	}

	args, err = p.parseNamed(args)
	if err != nil {
		return args, err
	}

	return args, err
}

func (p *ParserController) Validate() (err errors.ParseError) {
	for _, parser := range p.subparsers {
		err = parser.Validate()
		if err != nil {
			return err
		}
	}
	for _, parser := range p.positional {
		err = parser.Validate()
		if err != nil {
			return err
		}
	}
	for _, parser := range p.named {
		err = parser.Validate()
		if err != nil {
			return err
		}
	}
	return err
}

func (p *ParserController) parsePositional(args []string) (remainder []string, err errors.ParseError) {
	if len(args) == 0 {
		return args, nil
	}

	done := false
	prevArgsLen := len(args)
	prevFirstArgLen := len(args[0])

	// Positional: first positional is run until done, then second etc
	for _, parser := range p.positional {
		for !done {
			args, err = p.parseSub(args)
			if err != nil || len(args) == 0 {
				return args, err
			}
			args, err = parser.Parse(args)
			if err != nil || len(args) == 0 {
				return args, err
			}
			done = prevArgsLen == len(args) && prevFirstArgLen == len(args[0])
			prevArgsLen = len(args)
			prevFirstArgLen = len(args[0])
		}
	}

	return args, err
}

func (p *ParserController) parseNamed(args []string) (remainder []string, err errors.ParseError) {
	if len(args) == 0 {
		return args, nil
	}

	done := false
	prevArgsLen := len(args)
	prevFirstArgLen := len(args[0])

	// Named are run in cycle until done
	for !done {
		for _, parser := range p.named {
			args, err = p.parseSub(args)
			if err != nil || len(args) == 0 {
				return args, err
			}
			args, err = parser.Parse(args)
			if err != nil || len(args) == 0 {
				return args, err
			}
		}
		done = prevArgsLen == len(args) && prevFirstArgLen == len(args[0])
		prevArgsLen = len(args)
		prevFirstArgLen = len(args[0])
	}

	return args, err
}

func (p *ParserController) parseSub(args []string) (remainder []string, err errors.ParseError) {
	for _, subparser := range p.subparsers {
		args, err = subparser.Parse(args)
	}
	return args, err
}
