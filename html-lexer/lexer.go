package html_lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const eof = -1

type HtmlLexer struct {
	pos           int
	width         int
	start         int
	input         string
	items         chan Item
	startTagCount int
	endTagCount   int
	determiner    HTMLDeterminer
}

func (l *HtmlLexer) run() {
	for state := lexText; state != nil; {
		state = state(l)
	}
	close(l.items)
}

func (l *HtmlLexer) emit(t ItemType) {
	l.items <- Item{Type: t, Value: l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *HtmlLexer) Next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width =
		utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

func (l *HtmlLexer) Ignore() {
	l.start = l.pos
}

func (l *HtmlLexer) Backup() {
	l.pos -= l.width
}

func (l *HtmlLexer) Peek() rune {
	r := l.Next()
	l.Backup()
	return r
}

func (l *HtmlLexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.Next()) >= 0 {
		return true
	}
	l.Backup()
	return false
}

func (l *HtmlLexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.Next()) >= 0 {
	}
	l.Backup()
}

func (l *HtmlLexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- Item{
		Type:  ItemError,
		Value: fmt.Sprintf(format, args...),
	}
	return nil
}
