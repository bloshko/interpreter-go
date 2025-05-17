package parsing

type Scanner struct {
	Source  string
	Tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{Source: source, start: 0, current: 0, line: 1}
}

func (s *Scanner) ScanTokens() {
	for !s.isAtTheEnd() {
		s.start = s.current
		s.scanToken()
	}

	s.Tokens = append(s.Tokens, NewToken(EOF, "", nil, s.line))
}

func (s *Scanner) scanToken() {
	switch c := s.advance(); c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtTheEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case '\r':
	case '\t':
		break
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		// TODO ERROR
		break
	}

}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtTheEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtTheEnd() {
		// TODO ERROR unterminated string
		return
	}

	s.advance()

	value := s.Source[s.start+1 : s.current]

	// TODO check if works correctly
	// append token doesnt accept value
	s.appendToken(STRING, value)
}

func (s *Scanner) peek() rune {
	if s.isAtTheEnd() {
		return rune(0)
	}

	return rune(s.Source[s.current])
}

func (s *Scanner) isAtTheEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.appendToken(tokenType, nil)
}

func (s *Scanner) appendToken(tokenType TokenType, literal any) {
	text := s.Source[s.start : s.current+1]

	s.Tokens = append(s.Tokens, NewToken(tokenType, text, literal, s.line))
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtTheEnd() {
		return false
	}
	if rune(s.Source[s.current]) != expected {
		return false
	}
	s.current++

	return true
}

func (s *Scanner) advance() rune {
	character := rune(s.Source[s.current])
	s.current++

	return character
}
