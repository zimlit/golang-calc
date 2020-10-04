package scanner

import "errors"

type TokenType string

const (
	Number = "number"
	Plus   = "plus"
	Minus  = "minus"
	Star   = "star"
	Slash  = "slash"
	Lparen = "left paren"
	Rparen = "right paren"
	EOF    = "end of file"
)

type Token struct {
	tt     TokenType
	lexeme string
}

func newToken(tt TokenType, lexeme string) Token {
	return Token{
		tt,
		lexeme,
	}
}

func (t *Token) ToString() string {
	return string(t.tt) + " " + t.lexeme
}

type Scanner struct {
	source  string
	tokens  []Token
	current int
	start   int
}

func NewScanner(source string) Scanner {
	return Scanner{
		source,
		[]Token{},
		0,
		0,
	}
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	var err error = nil
	for !s.isAtEnd() {
		s.start = s.current
		e := s.scanToken()
		if e != nil {
			err = e
		}
	}

	s.tokens = append(s.tokens, newToken(EOF, ""))
	return s.tokens, err
}

func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() error {
	c := s.advance()
	switch c {
	case ' ':
	case '\r':
	case '\t':
	case '\n':
	case '(':
		s.addToken(Lparen)
	case ')':
		s.addToken(Rparen)
	case '-':
		s.addToken(Minus)
	case '+':
		s.addToken(Plus)
	case '/':
		s.addToken(Slash)
	case '*':
		s.addToken(Star)
	default:
		if isDigit(c) {
			s.number()
		} else {
			return errors.New("Unexpected char")
		}
	}
	return nil
}

func (s *Scanner) advance() rune {
	s.current++
	return charAt(s.source, s.current-1)
}

func charAt(str string, idx int) rune {
	for i, c := range str {
		if i == idx {
			return c
		}
	}
	return '\000'
}

func (s *Scanner) addToken(tt TokenType) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, newToken(tt, text))
}

func (s Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}
	return charAt(s.source, s.current)
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	s.addToken(Number)
}

func (s Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\000'
	}
	return charAt(s.source, s.current+1)
}
