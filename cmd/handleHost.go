package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hostDsn string

func init() {
	hostCmd.PersistentFlags().StringVarP(&domainDsn, "dsn", "d", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql")

	rootCmd.AddCommand(hostCmd)
	hostCmd.AddCommand(hostListCmd)
	hostCmd.AddCommand(hostAddCmd)

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

	Short: "List hosts of domain in database",
	Long:  `List all dynamic DNS host entries in the dynpower database of a specific domain. If a DSN is submitted by the flag --dsn, this DSN will be used. If no DSN is provided, dynpower-cli tries to use the environment variables DBHOST, DBUSER, DBNAME and DBPASSWORD.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command,
		args []string) {
		listHost(hostDsn, args[0])
	},
}

var hostAddCmd = &cobra.Command{
	Use:   "add [domain] [host]",
	Short: "Add host of domain to database",
	Long:  `Add host of domain to the dynpower database. The domain must already exist. If a DSN is submitted by the flag --dsn, this DSN will be used. If no DSN is provided, dynpower-cli tries to use the environment variables DBHOST, DBUSER, DBNAME and DBPASSWORD.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command,
		args []string) {
		addHost(hostDsn, args[0], args[1])
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
	//db := dbConn(dsn)
	//log.Println(dsn)
	//var maxStrlen int
	/*
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


		}*/
	fmt.Println("list hosts")

}

func addHost(dsn string, domain string, host string) {
	db := dbConn(dsn)
	//log.Println(dsn)
	//log.Println("in add...")
	//log.Println(domain)
	//log.Println(accessKey)

	// get domain id
	/*
		result, err := db.Prepare("INSERT INTO dynrecords (hostname, dt_created, dt_created) VALUES (?,now(), now())")
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
	*/
	fmt.Println("add host")
	defer db.Close()
}

/*
* Handle host-related commands
 */
func handleHost() {
	fmt.Println("\nUnknown or missing command.\nRun dynpower-cli domain --help to show available commands.")
}
