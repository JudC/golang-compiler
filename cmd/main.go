package main

import (
	"fmt"

	"github.com/JudC/golang-compiler/scanner"
)

func main() {
	t := scanner.Token{Kind: scanner.WHITESPACE, Val: "Whitespace"}

	fmt.Printf("%s", t.ToString())
}
