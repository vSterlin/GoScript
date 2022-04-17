package main

import (
	"fmt"

	"github.com/vSterlin/goscript/token"
)

func main() {

	// fmt.Printf("%#v\n %#v\n", first, second)

	// one := &parser.ParseResult[parser.Exp]{Result: &parser.NumberExp{1}}
	// two := &parser.ParseResult[parser.Exp]{Result: &parser.NumberExp{1}}

	// fmt.Printf("one.Equals(two): %v\n", one.Equals(two))

	// tokens := []token.Token{
	// 	&token.IfToken{},
	// 	&token.LeftParenToken{},
	// 	&token.NumberToken{1},
	// 	&token.EqualsToken{},
	// 	&token.NumberToken{1},
	// 	&token.RightParenToken{},
	// 	&token.LeftCurlyToken{},
	// 	&token.RightCurlyToken{},
	// 	&token.ElseToken{},
	// 	&token.LeftCurlyToken{},
	// 	&token.RightCurlyToken{},
	// }

	// tokens := []token.Token{
	// 	&token.LeftCurlyToken{},
	// 	&token.RightCurlyToken{},
	// }

	// tokens := []token.Token{&token.NumberToken{1}, &token.EqualsToken{}, &token.NumberToken{1}}

	tokenizer := token.NewTokenizer("fn int abcefg(lol int){}")
	tokens := tokenizer.Tokenize()

	for _, t := range tokens {
		fmt.Printf("%#v\n", t)
	}

	// fmt.Printf("tokens: %v\n", tokens)

}
