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
	t.Run("ScanTokens", func(t *testing.T) {
		source := "{>=12}"
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
	})

	t.Run("Comment", func(t *testing.T) {
		source := "//{}"

		scanner := NewScanner(source)
		scanner.ScanTokens()

		assert.Equal(t, scanner.Tokens, []Token{
			{
				Type:    EOF,
				Lexeme:  "",
				Literal: nil,
				Line:    1,
			},
		})

	})

	t.Run("New line", func(t *testing.T) {
		source := "//{}\n\n"

		scanner := NewScanner(source)
		scanner.ScanTokens()

		assert.Equal(t, scanner.line, 3)
	})

	t.Run("String", func(t *testing.T) {
		source := `"test"`

		scanner := NewScanner(source)
		scanner.ScanTokens()

		assert.Equal(t, scanner.Tokens, []Token{
			{
				Type:    STRING,
				Lexeme:  "\"test\"",
				Literal: "test",
				Line:    1,
			},
			{
				Type:    EOF,
				Lexeme:  "",
				Literal: nil,
				Line:    1,
			},
		})
	})

	t.Run("Number", func(t *testing.T) {
		source := "123"

		scanner := NewScanner(source)
		scanner.ScanTokens()

		assert.Equal(t, scanner.Tokens, []Token{
			{
				Type:    NUMBER,
				Lexeme:  "123",
				Literal: nil,
				Line:    1,
			},
			{
				Type:    EOF,
				Lexeme:  "",
				Literal: nil,
				Line:    1,
			},
		})
	})

	t.Run("Identifier", func(t *testing.T) {
		source := "int main {if}"

		scanner := NewScanner(source)
		scanner.ScanTokens()

		assert.Equal(t, scanner.Tokens, []Token{
			{
				Type:    IDENTIFIER,
				Lexeme:  "int",
				Literal: nil,
				Line:    1,
			},
			{
				Type:    IDENTIFIER,
				Lexeme:  "main",
				Literal: nil,
				Line:    1,
			},
			{
				Type:    LEFT_BRACE,
				Lexeme:  "{",
				Literal: nil,
				Line:    1,
			},
			{
				Type:    IF,
				Lexeme:  "if",
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
			},
		})
	})

}
