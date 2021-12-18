package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/a-poor/cypher/parser"
)

func main() {
	q := `MATCH (:Person {name: "Alice"})-[:KNOWS]->(:Person {name: "Bob"}) RETURN count(*)`

	// Setup the input
	is := antlr.NewInputStream(q)

	// Create the Lexer
	lexer := parser.NewCypherLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
