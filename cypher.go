package main

import "fmt"

func main() {

	query := `MATCH (n {name: "John"})-[:FRIEND]-(friend) WITH n, count(friend) AS friendsCount WHERE friendsCount > 3 RETURN n, friendsCount`
	fmt.Printf("Query: %s\n", query)

}
