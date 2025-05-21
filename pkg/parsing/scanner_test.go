package parsing

import "testing"

func TestNewScanner(t *testing.T) {
	source := "int main {}"
	scanner := *NewScanner(source)

	if !(scanner.Source == source && scanner.Tokens == nil && scanner.start == 0 && scanner.current == 0) {
		t.Errorf("Scanner created incorrectly")
	}
}
