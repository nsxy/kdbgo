package kdb_test

import (
	"fmt"

	"github.com/nsxy/kdbgo"
)

func ExampleDialKDB() {
	con, err := kdb.DialKDB("localhost", 5001, "")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}

	res, err := con.Call("til", kdb.Int(10))
	if err != nil {
		fmt.Println("Query failed:", err)
	}
	fmt.Println("Result:", res)
}
