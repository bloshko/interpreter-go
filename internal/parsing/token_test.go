package parsing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToken(t *testing.T) {
	token := NewToken(STRING, "", nil, 3)

	assert.Equal(t, token, &Token{
		Type:    STRING,
		Lexeme:  "",
		Literal: nil,
		Line:    3,
	})
}
