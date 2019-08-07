package html_lexer

type stateFn func(lexer *HtmlLexer) stateFn

func lexBlock(lexer *HtmlLexer) stateFn {
	for {
		if lexer.determiner.IsText(lexer) {
			return lexText(lexer)
		}

		if lexer.determiner.IsStartOpenTag(lexer) {
			
		}

		if lexer.Next() == eof { break }
	}
	lexer.emit(ItemEOF)
	return nil
}

func lexText(lexer *HtmlLexer) stateFn  {
	return nil
}

func lex()  {
	
}
