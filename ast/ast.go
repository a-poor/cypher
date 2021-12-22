package ast

type Node interface {
	// Text returns the the text of the node.
	Text() string
}

type Visitor interface {
}

type CypherQuery struct {
}
