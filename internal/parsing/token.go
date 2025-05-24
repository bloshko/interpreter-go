package parsing

import "fmt"

type TokenType int

const (
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR

	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

func (token TokenType) String() string {
	return fmt.Sprintf("%d", token)
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(Type TokenType, Lexeme string, Literal any, Line int) *Token {
	return &Token{Type, Lexeme, Literal, Line}
}

func (t *Token) String() string {
	return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
}
