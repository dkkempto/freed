package main

import (
	"fmt"

	"github.com/dkkempto/freed/parser/stl"
)

func main() {
	parser := stl.STLParser{}

	model := parser.ParseBinary("C:/development/personal/freed/example/dragon.stl")

	fmt.Println(model.String())
}
