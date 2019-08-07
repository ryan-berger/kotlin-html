package lexer

import (
	"bytes"
	"fmt"
	"strings"
)

type Node interface {
	Build() string
}

type Attribute struct {
	Key   string
	Value string
}

type HTMLNode struct {
	Element    string
	Children   []Node
	Attributes []Attribute
}


func (html *HTMLNode) Build() string {
	elemName := strings.TrimSpace(html.Element)
	b := bytes.NewBufferString("")
	b.WriteString(fmt.Sprintf("<%s>\n", elemName))
	for _, child := range html.Children {
		b.WriteString(child.Build())
	}
	b.WriteString(fmt.Sprintf("</%s>\n", elemName))
	return b.String()
}


type TextNode struct {
	Value string
}

func (text *TextNode) Build() string {
	return fmt.Sprintf("%s\n", text.Value)
}

type Parser struct {
	Nodes []Node
}

func (p *Parser) parse(input string) {
	_, items := Lex("lexer", input)

	for item := range items {
		if item.typ == itemQuote {
			p.Nodes = append(p.Nodes, p.parseQuoteText(item, items))
		}
		if item.typ == itemTag {
			p.Nodes = append(p.Nodes, p.parseHTMLText(item, items))
		}
	}
}

func (p *Parser) parseQuoteText(i item, itemStream chan item) *TextNode {
	textNode := &TextNode{}
	for item := range itemStream {
		if item.typ == itemQuoteText {
			textNode.Value = item.val
		}
		if item.typ == itemQuote {
			break
		}
	}
	return textNode
}

func (p *Parser) parseHTMLText(i item, itemStream chan item) *HTMLNode {
	htmlNode := &HTMLNode{
		Element: i.val,
	}
	for item := range itemStream {
		if item.typ == itemLeftParen {
			p.parseAttributes(itemStream)
		}
		if item.typ == itemTag {
			htmlNode.Children = append(htmlNode.Children, p.parseHTMLText(item, itemStream))
		}
		if item.typ == itemQuote {
			htmlNode.Children = append(htmlNode.Children, p.parseQuoteText(item, itemStream))
		}
		if item.typ == itemRightBrace {
			break
		}
	}
	return htmlNode
}

func (p *Parser) parseAttributes(itemStream chan item) []Attribute {
	var attributes []Attribute

	attr := Attribute{}

	for item := range itemStream {
		if item.typ == itemKey {
			attr.Key = item.val
		}
		if item.typ == itemValue {
			attr.Value = item.val
			attributes = append(attributes, attr)
			attr = Attribute{}
		}
		if item.typ == itemRightBrace {
			break
		}
	}

	return attributes
}

func (p *Parser) BuildString() string  {
	b := bytes.NewBufferString("")
	for _, node := range p.Nodes {
		b.WriteString(node.Build())
	}
	return b.String()
}


func NewParser(input string) *Parser {
	parser := &Parser{}
	parser.parse(input)
	return parser
}
