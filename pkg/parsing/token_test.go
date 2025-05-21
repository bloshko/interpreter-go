package parsing

import (
	"reflect"
	"testing"
)

func TestToken(t *testing.T) {
	token := NewToken(STRING, "", nil, 3)

	expected := Token{Type: STRING, Lexeme: "", Literal: nil, Line: 3}

	if reflect.DeepEqual(token, expected) {
		t.Errorf("%v does not equal %v", token, expected)
	}
}
