package lexer

import "fmt"

type itemType int

const (
	itemError itemType = iota

	itemEOF
	itemTag
	itemKey
	itemValue
	itemColon
	itemComma
	itemLeftBrace
	itemRightBrace
	itemLeftParen
	itemRightParen
	itemQuote
	itemQuoteText
)

type item struct {
	typ itemType
	val string
}

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val
	}
	if len(i.val) > 10 {
		return fmt.Sprintf("%.10q...", i.val)
	}
	return fmt.Sprintf("%q", i.val)
}

