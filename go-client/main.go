package main

import (
	"log"
        "os"
        "fmt"
	"github.com/gocql/gocql"
        "github.com/spf13/viper"
)

// Execute the CQL query
func executeQuery(session *gocql.Session, query string) error {
	return session.Query(query).Exec()
}

func retrieveQuery(session *gocql.Session, query string) error {
	// Execute the query
	iter := session.Query(query).Iter()

	// Get column names
	columns := iter.Columns()

	// Print column names
	fmt.Println("Columns:")
	for _, col := range columns {
		fmt.Println(col)
	}


        var id string
        var name string
        var metric string
        var host string

        for iter.Scan(&id, &name, &metric, &host) {
                fmt.Println("values:", id, name, metric, host)
        }

	// Check for any errors during iteration
	if err := iter.Close(); err != nil {
		return err
	}

	return nil
}

//example:
//go run . "select id,name,metric,host from collection_config;"

func main() {

        viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read configuration file: %s", err)
                return
	}

        host := viper.GetString("cassandra.host") + ":" + viper.GetString("cassandra.port")
        username := viper.GetString("cassandra.username")
	password := viper.GetString("cassandra.password")

	// Create a cluster configuration
	cluster := gocql.NewCluster(host)

	// Set the keyspace to use
	cluster.Keyspace = "telemetry" // Replace with your keyspace name

	// Set other options like username, password, etc. if required
        // Set authentication credentials
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}


	// Create a session
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// Perform database operations here
        query := os.Args[1] // Retrieve the query from command-line arguments
	if err := retrieveQuery(session, query); err != nil {
		log.Fatal(err)
	}
}


