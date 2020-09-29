package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn(dsn string) (db *sql.DB) {
	/*if
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWORD")

	db, err := sql.Open("mysql", dbuser+":"+dbpassword+"@tcp("+dbhost+":3306)/"+dbname)
	if err != nil {
		log.Println("Error by connecting database.")
		panic(err.Error())
	}
	return db*/
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

func List(dsn string) {
	db := dbConn(dsn)
	//log.Println(dsn)
	var maxStrlen int
	errLen := db.QueryRow("SELECT max(length(domainname)) as maxstrlen from domains").Scan(&maxStrlen)

	if errLen != nil {
		log.Println(errLen)
		os.Exit(1)
	}
	//	log.Println("max: ")
	//log.Println(maxStrlen)

	results, err := db.Query("SELECT domainname, dt_created, dt_updated FROM domains ORDER by domainname")

	if err != nil {
		log.Println("Database problem: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer db.Close()

	maxStrlen = maxStrlen + 2

	fmt.Println("Domains in database:")
	fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", "Domain", "Created", "Updated")
	var domainname, dtCreated, dtUpdated string
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&domainname, &dtCreated, &dtUpdated)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", domainname, dtCreated, dtUpdated)
	}

}

func add(dsn string, domain string, accessKey string) {
	db := dbConn(dsn)
	//log.Println(dsn)
	//log.Println("in add...")
	//log.Println(domain)
	//log.Println(accessKey)

	hashedAccessKey, err := HashPassword(accessKey)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//log.Println(hashedAccessKey)

	result, err := db.Prepare("INSERT INTO domains(domainname, access_key, dt_created) VALUES(?,?, now())")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	_, insertErr := result.Exec(domain, hashedAccessKey)
	if insertErr != nil {
		fmt.Println(insertErr.Error())
		os.Exit(1)
	}

	fmt.Println("Domain " + domain + " added successfully")

	defer db.Close()
}

/*
* Handle domain-related commands
 */
func HandleDomain() {
	//	func HandleDomain(fs *flag.FlagSet, dsn string) {
	fmt.Println("handle Domain command")
	//fmt.Println("DSN: " + dsn)

	//fmt.Println(fs.Args())
	/*switch fs.Arg(0) {
	case "list":
		//fmt.Println("Command: list")
		list(dsn)
	case "add":
		//fmt.Println("Command: add")
		domain := fs.Arg(1)
		if len(domain) < 1 {
			fmt.Println("\nDomain parameter missing. \n")
			os.Exit(1)
		}
		accessKey := fs.Arg(2)
		if len(accessKey) < 1 {
			fmt.Println("\nAccess Key parameter missing. \n")
			os.Exit(1)
		}
		add(dsn, domain, accessKey)
	default:
		fmt.Println("Unknown command.")
	}*/

}
