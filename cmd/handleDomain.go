package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn(dsn string) (db *sql.DB) {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error by connecting database.")
		panic(err.Error())
	}
	return db
}

/*
 *	Check database connection by performing a query, exit in error case
 */
func checkDb(dsn string) {
	db := dbConn(dsn)
	_, err := db.Query("SELECT r.hostname, d.domainname, d.access_key FROM dynrecords r, domains d WHERE d.id=r.domain_id")
	if err != nil {
		log.Println("Database problem: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}
	return
}

func list(dsn string) {
	db := dbConn(dsn)
	log.Println(dsn)
	results, err := db.Query("SELECT domainname FROM domains ORDER by domainname")

	if err != nil {
		log.Println("Database problem: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer results.Close()

	var content string
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&content)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute

		log.Printf(content)
	}

}

/*
* Handle domain-related commands
 */
func HandleDomain(fs *flag.FlagSet, dsn string) {
	fmt.Println("handle Domain command")

	fmt.Println(fs.Args())
	switch fs.Arg(0) {
	case "list":
		fmt.Println("Command: list")
		list(dsn)
	case "add":
		fmt.Println("Command: add")
	default:
		fmt.Println("Unknown command.")
	}

}
