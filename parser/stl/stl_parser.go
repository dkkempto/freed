package stl

import (
	"github.com/dkkempto/freed/parser"
)

type STLParser struct {
}

func (p STLParser) parse(path string) []parser.Model {
	return make([]parser.Model, 0)
}

func main() {
	p := STLParser{}
	p.parse("./res/sub2pewdiepie.stl")
}
