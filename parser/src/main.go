package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

type Lexer struct {
	input string
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

type Parser struct {
	lexer *Lexer
}
