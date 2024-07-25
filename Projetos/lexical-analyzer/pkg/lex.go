package pkg

import (
	"fmt"
	"regexp"
	"strings"
)

type TokenType string

const (
	INT       TokenType = "INT"
	STRING    TokenType = "STRING"
	VAR       TokenType = "VAR"
	EQ        TokenType = "EQ"
	NUM       TokenType = "NUM"
	ADD       TokenType = "ADD"
	SUB       TokenType = "SUB"
	MUL       TokenType = "MUL"
	GT        TokenType = "GT"
	LT        TokenType = "LT"
	SEMICOLON TokenType = "SEMICOLON"
	CONST     TokenType = "CONST"
	ERROR     TokenType = "ERROR"
)

type Token struct {
	Type   TokenType
	Lexeme string
}

var tokenDefinitions = []struct {
	Type  TokenType
	Regex *regexp.Regexp
}{
	{INT, regexp.MustCompile(`^int\b`)},
	{STRING, regexp.MustCompile(`^string\b`)},
	{VAR, regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*`)},
	{EQ, regexp.MustCompile(`^=`)},
	{NUM, regexp.MustCompile(`^[0-9]+`)},
	{ADD, regexp.MustCompile(`^\+`)},
	{SUB, regexp.MustCompile(`^-`)},
	{MUL, regexp.MustCompile(`^\*`)},
	{GT, regexp.MustCompile(`^>`)},
	{LT, regexp.MustCompile(`^<`)},
	{SEMICOLON, regexp.MustCompile(`^;`)},
	{CONST, regexp.MustCompile(`^"[^"]*"`)},
}

func Lex(input string) ([]Token, error) {
	tokens := []Token{}
	remaining := strings.TrimSpace(input)

	for len(remaining) > 0 {
		matched := false
		for _, def := range tokenDefinitions {
			re := def.Regex
			loc := re.FindStringIndex(remaining)
			if loc != nil && loc[0] == 0 {
				lexeme := remaining[:loc[1]]
				if def.Type != ERROR {
					tokens = append(tokens, Token{Type: def.Type, Lexeme: lexeme})
				}
				remaining = remaining[loc[1]:]
				matched = true
				break
			}
		}
		if !matched {
			return nil, fmt.Errorf("erro léxico: '%s'", remaining)
		}
		remaining = strings.TrimSpace(remaining)
	}

	return tokens, nil
}
