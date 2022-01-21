package main

import (
	"bytes"
	"log"
	"os"

	"golang.org/x/exp/ebnf"
)

const ebnfPath = "cypher.ebnf"

func main() {
	//query := `MATCH (n {name: "John"})-[:FRIEND]-(friend) WITH n, count(friend) AS friendsCount WHERE friendsCount > 3 RETURN n, friendsCount`
	//fmt.Printf("Query: %s\n", query)

	in, err := os.ReadFile(ebnfPath)
	if err != nil {
		log.Panic(err)
	}

	gram, err := ebnf.Parse(ebnfPath, bytes.NewReader(in))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Grammar length: %d\n", len(gram))

}
