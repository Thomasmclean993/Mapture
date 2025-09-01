package parser

// Registry maps supported config sources to their parser implementations.
var Registry = map[string]Parser{
	"aerospace": AerospaceParser{},
	"nvim":      NvimParser{},
}
