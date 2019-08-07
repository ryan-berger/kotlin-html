package html_lexer

type ItemType int

const (
	ItemEOF ItemType = iota
	ItemError

	ItemTag
	ItemAttributeKey
	ItemAttributeValue
	ItemText
)

type Item struct {
	Type  ItemType
	Value string
}
