package parsing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewScanner(t *testing.T) {
	source := "int main {}"
	scanner := NewScanner(source)

	assert.Equal(t, scanner, &Scanner{
		Source:  source,
		Tokens:  nil,
		start:   0,
		current: 0,
		line:    1,
	})

}

func TestScanTokens(t *testing.T) {
	source := "int main {>=12}"
	scanner := NewScanner(source)
	scanner.ScanTokens()

	assert.Equal(t, scanner.Tokens, []Token{{
		Type:    LEFT_BRACE,
		Lexeme:  "{",
		Literal: nil,
		Line:    1,
	},
		{
			Type:    GREATER_EQUAL,
			Lexeme:  ">=",
			Literal: nil,
			Line:    1,
		},
		{
			Type:    NUMBER,
			Lexeme:  "12",
			Literal: nil,
			Line:    1,
		},
		{
			Type:    RIGHT_BRACE,
			Lexeme:  "}",
			Literal: nil,
			Line:    1,
		},
		{
			Type:    EOF,
			Lexeme:  "",
			Literal: nil,
			Line:    1,
		}})

}
