package main

import (
	"io"
	"os"

	"github.com/alecthomas/participle/v2/lexer"
)

const filepath = "cypher.ebnf"

func readFile() string {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return string(b)
}

type EBNF struct {
	Productions []*Production `@@*`
}

type Comment struct{} // TODO - For now, assume comments have been stripped out

type Production struct {
	Production string      `@Ident "="`
	Expression *Expression `@@ "."`
}

type Expression struct {
	Alternatives []*Sequence `@@ ( "|" @@ )*`
}

type SubExpression struct{}

type Sequence struct{}

type Term struct {
	Negation bool           `@("~")?`
	Name     string         `(   @Ident`
	Literal  string         `  | @String`
	Token    string         `  | "<" @Ident ">"`
	Group    *SubExpression `  | @@ )`

	Repetition string `@("*" | "+" | "?" | "!")?`
}

func main() {
	// Read in the file
	f := readFile()

	// Create a lexer
	lex := lexer.MustStateful(lexer.Rules{})

	// Parse the file

}
