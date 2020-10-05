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
var domainForce bool

func init() {
	domainCmd.PersistentFlags().StringVarP(&domainDsn, "dsn", "d", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql")
	domainRemoveCmd.Flags().BoolVarP(&domainForce, "force", "f", false, "If true, delete domain with all hosts. If false (default), a domain isn't deleted if any host of the domain exists.")

	rootCmd.AddCommand(domainCmd)
	domainCmd.AddCommand(domainListCmd)
	domainCmd.AddCommand(domainAddCmd)
	domainCmd.AddCommand(domainRemoveCmd)

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

var domainRemoveCmd = &cobra.Command{
	Use:   "remove [domain]",
	Short: "Remove domain from database",
	Long:  `Remove domain from dynpower database. If a DSN is submitted by the flag --dsn, this DSN will be used. If no DSN is provided, dynpower-cli tries to use the environment variables DBHOST, DBUSER, DBNAME and DBPASSWORD.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command,
		args []string) {
		removeDomain(domainDsn, args[0])
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

	var maxStrlenSQL sql.NullInt64
	var maxStrlen int

	errLen := db.QueryRow("SELECT max(length(domainname)) as maxstrlen from domains").Scan(&maxStrlenSQL)

	if errLen != nil {
		fmt.Println(errLen)
		os.Exit(1)
	}
	if maxStrlenSQL.Valid {
		maxStrlen = int(maxStrlenSQL.Int64)
	} else {
		maxStrlen = 0
	}
	if maxStrlen == 0 {
		fmt.Println("No domain in database, use add command to create domain entry")
		os.Exit(0)
	}

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

func removeHosts(dsn string, domain string) {
	db := dbConn(dsn)
	defer db.Close()
	//log.Println(dsn)
	//log.Println("in add...")
	//log.Println(domain)
	//log.Println(accessKey)

	var domainID int
	// todo: check existence of domain
	// remove host

	// get domain id
	errLen := db.QueryRow("SELECT id from domains d WHERE d.domainname=?", domain).Scan(&domainID)

	if errLen != nil {
		fmt.Println(errLen)
		os.Exit(1)
	}

	result, err := db.Prepare("DELETE FROM dynrecords WHERE domain_id=?")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	deleteResult, removeErr := result.Exec(domainID)
	if removeErr != nil {
		fmt.Println(removeErr.Error())
		os.Exit(1)
	}

	rowsAffected, removeErr := deleteResult.RowsAffected()
	if removeErr != nil {
		fmt.Println(removeErr.Error())
		os.Exit(1)
	}
	if rowsAffected == 0 {
		fmt.Println("No host entry deleted")
	} else {
		fmt.Println("Hosts of domain " + domain + " removed successfully")
	}

}

func removeDomain(dsn string, domain string) {
	db := dbConn(dsn)

	var count int
	err := db.QueryRow("SELECT count(*) as cnt FROM dynrecords r, domains d WHERE d.id=r.domain_id AND d.domainname=?", domain).Scan(&count)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer db.Close()
	if count > 0 && domainForce == true {

		fmt.Println("Delete host entries...")
		removeHosts(dsn, domain)

	} else if count > 0 && domainForce == false {
		fmt.Println("Could not delete domain, because there are host entries of domain " + domain + ".\nPlease delete host entries first or use --force flag.")
		os.Exit(0)
	}
	// count = 0
	// delete domain entry

	result, err := db.Prepare("DELETE FROM domains WHERE domainname=?")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	deleteResult, removeErr := result.Exec(domain)
	if removeErr != nil {
		fmt.Println(removeErr.Error())
		os.Exit(1)
	}

	rowsAffected, removeErr := deleteResult.RowsAffected()
	if removeErr != nil {
		fmt.Println(removeErr.Error())
		os.Exit(1)
	}
	if rowsAffected == 0 {
		fmt.Println("Nothing deleted, does the domain " + domain + " exist?")
	} else {
		fmt.Println("Domain " + domain + " removed successfully")
	}

}

/*
* Handle domain-related commands
 */
func handleDomain() {
	fmt.Println("\nUnknown or missing command.\nRun dynpower-cli domain --help to show available commands.")
}
