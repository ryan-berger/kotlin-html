package html_lexer

type Taper interface {
	Next() rune
	Peek() rune
	Ignore()
}
