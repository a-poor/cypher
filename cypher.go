package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/a-poor/cypher/parser"
)

type cypherListener struct {
	*parser.BaseCypherListener
}

func (l *cypherListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("Entering %q %T | %+v\n", ctx.GetText(), ctx, ctx)
}

func (l *cypherListener) VisitErrorNode(node antlr.ErrorNode) {
	panic(fmt.Errorf("%q, %T, %+v", node.GetText(), node, node))
}

func main() {
	// q := `MATCH (:Person {name: "Alice"})-[:KNOWS]-<(:Person {name: "Bob"}) RETURN count(*)`
	q := `CREATE (n:Person {name: 'Andy', title: 'Developer'})`

	// Setup the input
	is := antlr.NewInputStream(q)

	// Create the Lexer
	lexer := parser.NewCypherLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCypherParser(stream)

	// Finally parse the expression (by walking the tree)
	// The "listener" is the action triggered when visiting each node
	listener := &cypherListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.OC_Cypher())
}
