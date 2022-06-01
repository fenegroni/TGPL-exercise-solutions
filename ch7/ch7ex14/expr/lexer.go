package expr

import (
	"fmt"
	"text/scanner"
)

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// ---- lexer ----

// This lexer is similar to the one described in Chapter 13.
type lexer struct {
	scan  scanner.Scanner
	token rune // current lookahead token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

type lexPanic string

// describe returns a string describing the current token, for use in errors.
func (lex *lexer) describe() string {
	switch lex.token {
	case scanner.EOF:
		return "end of file"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", lex.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", lex.text())
	}
	return fmt.Sprintf("%q", lex.token) // any other rune
}
