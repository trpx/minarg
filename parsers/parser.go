package parsers

type Parser interface {
	SubParser()
	Int()
	String()
	IntSlice()
	StringSlice()
}
