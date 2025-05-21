package parsing

import (
	"testing"
)

func TestToken(t *testing.T) {
	token := *NewToken(STRING, "", nil, 3)

	if !(token.Type == STRING && token.Line == 3 && token.Lexeme == "" && token.Literal == nil) {
		t.Errorf("Token is generated incorrectly")
	}
}
