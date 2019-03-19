package main

import (
	"fmt"

	"github.com/dkkempto/freed/parser/stl"
)

func main() {
	parser := stl.STLParser{}

	models := parser.Parse("C:/development/personal/freed/example/Star.stl")

	for _, model := range models {
		fmt.Println(model.String())
	}
}
