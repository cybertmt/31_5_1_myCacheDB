package main

import (
	"fmt"
	"mycachedb/pkg/mycachedb"
)

func main() {
	db := mycachedb.New()
	db.SET("a", "test_b")
	fmt.Printf("('a': %s), ('b': %s))", db.GET("a"), db.GET("b"))
}
