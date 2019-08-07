package main

import (
	"fmt"
	"github.com/ryan-berger/kotlin-html/lexer"
)

const TEST = `
html(id: "hi") {
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p { "I love you!" }
    p {
        "I love you!"
        p { "I love you!" }
        p { "I love you!" }
        p { "I love you!" }
        p { "I love you!" }
        p { "I love you!" }
        p {
            "I love you!"
            p { "I love you!" }
            p { "I love you!" }
        }
    }
}
`

func main()  {
	p := lexer.NewParser(TEST)
	fmt.Println(p.BuildString())
}


