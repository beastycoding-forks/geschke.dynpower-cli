package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var hostDsn string

func init() {
	hostCmd.PersistentFlags().StringVarP(&domainDsn, "dsn", "d", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql")

	rootCmd.AddCommand(hostCmd)
	hostCmd.AddCommand(hostListCmd)
	hostCmd.AddCommand(hostAddCmd)
	hostCmd.AddCommand(hostRemoveCmd)

}

var hostCmd = &cobra.Command{
	Use: "host",

	Short: "Manage host entries",
	Long:  `Manage dynpower host entries of a specific domain in database.`,
	Run: func(cmd *cobra.Command,
		args []string) {
		handleHost()
	},
}

var hostListCmd = &cobra.Command{
	Use: "list [domain]",

	Short: "List hosts of the domain in database",
	Long:  `List all dynamic DNS host entries in the dynpower database of a specific domain. If a DSN is submitted by the flag --dsn, this DSN will be used. If no DSN is provided, dynpower-cli tries to use the environment variables DBHOST, DBUSER, DBNAME and DBPASSWORD.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command,
		args []string) {
		listHost(hostDsn, args[0])
	},
}

var hostAddCmd = &cobra.Command{
	Use:   "add [domain] [host]",
	Short: "Add host of the domain to database",
	Long:  `Add host of the domain to the dynpower database. The domain must already exist. If a DSN is submitted by the flag --dsn, this DSN will be used. If no DSN is provided, dynpower-cli tries to use the environment variables DBHOST, DBUSER, DBNAME and DBPASSWORD.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command,
		args []string) {
		addHost(hostDsn, args[0], args[1])
	},
}

var hostRemoveCmd = &cobra.Command{
	Use:   "remove [domain] [host]",
	Short: "Remove host of the domain from database",
	Long:  `Remove host of the domain from the dynpower database. The domain must already exist. If a DSN is submitted by the flag --dsn, this DSN will be used. If no DSN is provided, dynpower-cli tries to use the environment variables DBHOST, DBUSER, DBNAME and DBPASSWORD.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command,
		args []string) {
		removeHost(hostDsn, args[0], args[1])
	},
}

/*func dbConn(dsn string) (db *sql.DB) {
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
}*/

func listHost(dsn string, domain string) {
	db := dbConn(dsn)

	var maxStrlenSQL sql.NullInt64
	var maxStrlen int

	errLen := db.QueryRow("SELECT max(length(r.hostname)) as maxstrlen from dynrecords r, domains d WHERE r.domain_id=d.id AND d.domainname=?", domain).Scan(&maxStrlenSQL)

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
		fmt.Printf("No host entry for domain %s found.\n", domain)
		os.Exit(0)
	}

	maxStrlen = len(domain) + 1 + maxStrlen + 2

	results, err := db.Query("SELECT d.domainname, r.hostname, r.dt_created, r.dt_updated FROM domains d, dynrecords r WHERE d.domainname=? AND d.id=r.domain_id ORDER by r.hostname", domain)

	if err != nil {
		log.Println("Database problem: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer db.Close()

	fmt.Println("Hosts in database:")
	fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", "Host", "Created", "Updated")
	var domainname, host, dtCreated, dtUpdated string
	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&domainname, &host, &dtCreated, &dtUpdated)
		if err != nil {
			fmt.Println(err.Error()) // proper error handling instead of panic in your app
			os.Exit(1)
		}
		// and then print out the tag's Name attribute
		fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", host+"."+domainname, dtCreated, dtUpdated)

	}

}

func addHost(dsn string, domain string, host string) {
	db := dbConn(dsn)
	defer db.Close()
	//log.Println(dsn)
	//log.Println("in add...")
	//log.Println(domain)
	//log.Println(accessKey)

	var domainId int
	// todo: check existence of domain
	// add host

	// get domain id
	errLen := db.QueryRow("SELECT id from domains d WHERE d.domainname=?", domain).Scan(&domainId)

	if errLen != nil {
		fmt.Println(errLen)
		os.Exit(1)
	}

	fmt.Println("domain id: ")
	fmt.Println(domainId)

	result, err := db.Prepare("INSERT INTO dynrecords (domain_id, hostname, dt_created, dt_updated) VALUES (?,?,now(), now())")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	_, insertErr := result.Exec(domainId, host)
	if insertErr != nil {
		fmt.Println(insertErr.Error())
		os.Exit(1)
	}

	fmt.Println("Host " + host + "." + domain + " added successfully")

}

func removeHost(dsn string, domain string, host string) {
	db := dbConn(dsn)
	defer db.Close()
	//log.Println(dsn)
	//log.Println("in add...")
	//log.Println(domain)
	//log.Println(accessKey)

	var domainId int
	// todo: check existence of domain
	// remove host

	// get domain id
	errLen := db.QueryRow("SELECT id from domains d WHERE d.domainname=?", domain).Scan(&domainId)

	if errLen != nil {
		fmt.Println(errLen)
		os.Exit(1)
	}

	result, err := db.Prepare("DELETE FROM dynrecords WHERE domain_id=? AND hostname=?")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	deleteResult, removeErr := result.Exec(domainId, host)
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
		fmt.Println("Nothing deleted, does host " + host + "." + domain + " exist?")
	} else {
		fmt.Println("Host " + host + "." + domain + " removed successfully")

	}

}

/*
* Handle host-related commands
 */
func handleHost() {
	fmt.Println("\nUnknown or missing command.\nRun dynpower-cli domain --help to show available commands.")
}
