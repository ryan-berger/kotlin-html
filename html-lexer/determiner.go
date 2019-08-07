package html_lexer

// Determine what part of the lexer to use

type HTMLDeterminer interface {
	IsText(Taper) bool
	IsStartOpenTag(Taper) bool
	IsEndOpenTag(Taper) bool
	IsStartCloseTag(Taper) bool
	IsEndCloseTag(Taper) bool
}
