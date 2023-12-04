package main

import (
	"fmt"
	"grupie-tracker/internal/parser"
)

func main() {

	model, err := parser.Parse()

	if err != nil {
		return
	}
	fmt.Println(model)
}
