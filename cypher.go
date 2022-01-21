package cypher

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/olekukonko/tablewriter"

	"github.com/a-poor/cypher/parser"
)

func Parse(query string) *parser.CypherParser {
	// Setup the input
	is := antlr.NewInputStream(query)

	// Create the Lexer
	lexer := parser.NewCypherLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCypherParser(stream)

	return p
}

type cypherListener struct {
	*parser.BaseCypherListener

	tw *tablewriter.Table
}

func (l *cypherListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	l.tw.Append([]string{
		fmt.Sprintf("%T", ctx),
		ctx.GetText(),
	})

	// fmt.Printf("Entering %q | %T | %+v\n", ctx.GetText(), ctx, ctx)
}

func (l *cypherListener) VisitErrorNode(node antlr.ErrorNode) {
	panic(fmt.Errorf("%q, %T, %+v", node.GetText(), node, node))
}

/*
func main() {
	// q := `MATCH (:Person {name: "Alice"})-[:KNOWS]->(:Person {name: "Bob"}) RETURN count(*)`
	// q := `CREATE (n:Person {name: 'Andy', title: 'Developer'})`
	q := `MATCH (p1:Person {name: "Alice", IsCool: $isCool})-[:KNOWS]->(p2:Person) WHERE p2.name != "Tom" RETURN count(*)`

	// Setup the input
	is := antlr.NewInputStream(q)

	// Create the Lexer
	lexer := parser.NewCypherLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCypherParser(stream)

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetColWidth(60)
	tw.SetHeader([]string{"Type", "Text"})

	// Finally parse the expression (by walking the tree)
	// The "listener" is the action triggered when visiting each node
	listener := &cypherListener{tw: tw}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.OC_Cypher())

	tw.Render()
}

*/
