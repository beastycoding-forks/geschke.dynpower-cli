package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var domainDsn string

func init() {
	domainCmd.PersistentFlags().StringVarP(&domainDsn, "dsn", "d", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql")

	rootCmd.AddCommand(domainCmd)
	domainCmd.AddCommand(domainListCmd)
	domainCmd.AddCommand(domainAddCmd)

}

var domainCmd = &cobra.Command{
	Use: "domain",

	Short: "Manage domain entries",
	Long:  `Manage dynpower domain entries in database.`,
	Run: func(cmd *cobra.Command,
		args []string) {
		handleDomain()
	},
}

var domainListCmd = &cobra.Command{
	Use: "list",

	Short: "List domains in database",
	Long:  `List all domains in the dynpower database. If a DSN is submitted by the flag --dsn, this DSN will be used. If no DSN is provided, dynpower-cli tries to use the environment variables DBHOST, DBUSER, DBNAME and DBPASSWORD.`,
	Run: func(cmd *cobra.Command,
		args []string) {
		listDomain(domainDsn)
	},
}

var domainAddCmd = &cobra.Command{
	Use:   "add [domain] [access key]",
	Short: "Add domain with access key to database",
	Long:  `Add domain with access key to dynpower database. If a DSN is submitted by the flag --dsn, this DSN will be used. If no DSN is provided, dynpower-cli tries to use the environment variables DBHOST, DBUSER, DBNAME and DBPASSWORD.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command,
		args []string) {
		addDomain(domainDsn, args[0], args[1])
	},
}

func dbConn(dsn string) (db *sql.DB) {
	if len(dsn) < 1 {
		dbname := os.Getenv("DBNAME") // maybe use LookupEnv to detect if env variable exists
		dbhost := os.Getenv("DBHOST")
		dbuser := os.Getenv("DBUSER")
		dbpassword := os.Getenv("DBPASSWORD")
		dbport := os.Getenv("DBPORT")
		if len(dbport) < 1 {
			dbport = "3306"
		}
		if len(dbname) >= 1 && len(dbhost) >= 1 && len(dbuser) >= 1 && len(dbpassword) >= 1 {
			dsn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname

		} else {
			fmt.Println("No database connect parameter found, exiting. Please use --dsn or environment variables to define database connection.")
			os.Exit(1)
		}

	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error by connecting database.")
		fmt.Println(err.Error())
		os.Exit(1)
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

func listDomain(dsn string) {
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

func addDomain(dsn string, domain string, accessKey string) {
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
func handleDomain() {
	fmt.Println("\nUnknown or missing command.\nRun dynpower-cli domain --help to show available commands.")
}
