package main

import (
	"fmt"

	"github.com/gocql/gocql"
)

func main() {
	cluster := gocql.NewCluster("10.233.8.211")
	cluster.Port = 31631
	cluster.Keyspace = "my_keyspace"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer session.Close()

        // Example query execution
	query := session.Query("SELECT * FROM my_table")

	// Execute the query
	iter := query.Iter()
	defer iter.Close()

        var id string
        var age int
	var name string

	for iter.Scan(&id,&age, &name) {
		fmt.Println("ID:", id, "Name:", name, "Age:",age)
	}

	if err := iter.Close(); err != nil {
		fmt.Println("Error:", err)
	}

 }

