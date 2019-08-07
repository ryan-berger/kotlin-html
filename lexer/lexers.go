package lexer

import (
	"strings"
	"unicode"
)

const LParen = "("
const RParen = ")"
const LBrace = "{"
const RBrace = "}"
const Colon = ":"
const Comma = ","
const Quote = `"`

func lexText(lex *lexer) stateFn {
	for {
		if strings.HasPrefix(lex.input[lex.pos:], Quote) {
			return lexLQuote
		}

		if strings.HasPrefix(lex.input[lex.pos:], RBrace) {
			return lexRightBrace
		}

		next := lex.next()

		if next == eof { break }


		switch {
		case unicode.IsSpace(next):
			lex.ignore()
		case isAlphaNumeric(next):
			return lexIdentifier
		default:
			return lex.errorf("expected identifier or rparen but received %s", string(next))
		}
	}
	lex.emit(itemEOF)
	return nil
}

// identifier is
// <tag> <lparen> <key:value,> <rparen> OR
// <tag> <lbrace>
func lexIdentifier(lex *lexer) stateFn {
	for {
		if strings.HasPrefix(lex.input[lex.pos:], LParen) {
			lex.emit(itemTag)
			return lexLeftParen
		}

		if strings.HasPrefix(lex.input[lex.pos:], LBrace) {
			lex.emit(itemTag)
			return lexLeftBrace
		}

		next := lex.next()
		switch {
		case isAlphaNumeric(next) || unicode.IsSpace(next):
			continue
		case next == eof:
			return lex.errorf("found")
		default:
			return lex.errorf("expected ( or {, found %s", string(next))

		}
	}
}


// move down and look for attribute
func lexLeftParen(lex *lexer) stateFn  {
	lex.pos += len(LParen)
	lex.emit(itemLeftParen)
	return lexAttribute
}

// move down and look for next identifier
func lexLeftBrace(lex *lexer) stateFn  {
	lex.pos += len(LBrace)
	lex.emit(itemLeftBrace)
	lex.lBraceCount++
	return lexText
}

// look for attribute
func lexAttribute(lex *lexer) stateFn {
	for {
		if strings.HasPrefix(lex.input[lex.pos:], Colon) {
			lex.emit(itemKey)
			return lexColon
		}

		next := lex.next()
		if unicode.IsSpace(next) || unicode.IsLetter(next) || unicode.IsNumber(next) {
			continue
		}

		if next == eof {
			lex.emit(itemEOF)
			return lex.errorf("found eof when expecting attribute")
		}
	}
}

func lexColon(lex *lexer) stateFn {
	lex.pos += len(Colon)
	lex.emit(itemColon)
	return lexValue
}

func lexValue(lex *lexer) stateFn  {
	for  {
		if strings.HasPrefix(lex.input[lex.pos:], Comma) {
			lex.emit(itemValue)
			return lexComma
		}

		if strings.HasPrefix(lex.input[lex.pos:], RParen) {
			lex.emit(itemValue)
			return lexRightParen
		}

		if lex.next() == eof {
			lex.errorf("expected comma or rparen, found eof")
		}
	}
}

func lexComma(lex *lexer) stateFn  {
	lex.pos += len(Comma)
	lex.emit(itemComma)
	return lexAttribute
}

func lexRightParen(lex *lexer) stateFn  {
	lex.pos += len(RParen)
	lex.emit(itemRightParen)

	for {
		if strings.HasPrefix(lex.input[lex.pos:], LBrace) {
			return lexLeftBrace
		}

		if !unicode.IsSpace(lex.next()) {
			lex.errorf("missing {")
		}
	}

}

func lexRightBrace(lex *lexer) stateFn {
	lex.pos += len(RBrace)
	lex.emit(itemRightBrace)
	lex.rBraceCount++
	if lex.rBraceCount > lex.lBraceCount {
		return lex.errorf("more right braces than left")
	}
	return lexText
}

func lexLQuote(lex *lexer) stateFn  {
	lex.pos += len(Quote)
	lex.emit(itemQuote)
	return lexQuoteText
}

func lexQuoteText(lex *lexer) stateFn  {
	for {
		if lex.next() == eof { return lex.errorf("expecting end quote, got eof")}
		if strings.HasPrefix(lex.input[lex.pos:], Quote) {
			lex.emit(itemQuoteText)
			return lexRQuote
		}
	}
}

func lexRQuote(lex *lexer) stateFn {
	lex.pos += len(Quote)
	lex.emit(itemQuote)
	return lexText
}