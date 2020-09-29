package main

import (
	"github.com/geschke/dynpower-cli/cmd"
)

func main() {
	cmd.Exec()

	//var dsn string

	//flag.StringVar(&dsn, "dsn", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql#dsn-data-source-name")
	/*
		hostCmd := flag.NewFlagSet("host", flag.ContinueOnError)
		domainCmd := flag.NewFlagSet("domain", flag.ContinueOnError)

		// todo: evaluate cobra library or something else as replacement
		hostDsn := hostCmd.String("dsn", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql#dsn-data-source-name")
		domainDsn := domainCmd.String("dsn", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql#dsn-data-source-name")

		domainCmd.Parse(os.Args[2:])
		log.Println("hostDsn ist: ")
		log.Println(*hostDsn)

		log.Println("domainDsn ist: ")
		log.Println(*domainDsn)

	*/
	/*dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWORD")
	if len(dbname) >= 1 && len(dbhost) >= 1 && len(dbuser) >= 1 && len(dbpassword) >= 1 {
		dsn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":3306)/" + dbname

	}
	if len(*hostDsn) > 0 {
		dsn = *hostDsn
	}
	if len(*domainDsn) > 0 {
		dsn = *domainDsn
	}

	log.Println(*domainDsn)
	log.Println("dsn: " + dsn)
	*/

	/*dbPasswordPtr := flag.String("password", "", "Database password")
	dbHostPtr := flag.String("host", "", "Database server")
	dbNamePtr := flag.String("dbname")
	*/
	//numbPtr := flag.Int("numb", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")

	//var svar string
	//flag.StringVar(&svar, "svar", "bar", "a string var")

	// todo maybe: use flag subcommands

	/*
		flag.Parse()
		log.Println("hostDsn ist: ")
		log.Println(*hostDsn)

		log.Println("domainDsn ist: ")
		log.Println(*domainDsn)

		switch flag.Arg(0) {
		case "encrypt":
			password := flag.Arg(1)
			if len(password) < 1 {
				fmt.Println("\nPassword parameter missing. \n")
				os.Exit(1)
				//panic(err.Error()) // proper error handling instead of panic in your app
			}
			cmd.Encrypt(password)
		case "domain":
			command := flag.Arg(1)
			log.Println(command)
			if len(command) < 1 {
				fmt.Println("\nManaga domains.\n")
				fmt.Println("Available commands:")
				fmt.Println("\tlist\t\t List domains in database.")
				fmt.Println("\tadd <domain> <access key>\t Add domain with access key to database.\n")

				os.Exit(0)
			}
			domainCmd.Parse(os.Args[2:])
			fmt.Println("Execute subcommand 'domain'...")
			fmt.Println("  dsn:", *domainDsn)
			fmt.Println("  tail:", domainCmd.Args())

			cmd.HandleDomain(domainCmd, dsn)

			fmt.Println("")
		case "host":
			command := flag.Arg(1)
			if len(command) < 1 {
				fmt.Println("\nManage hosts.\n")
				fmt.Println("Available commands:")
				fmt.Println("\tlist <domain>\t List hosts of <domain> in database.")
				fmt.Println("\tadd <domain> <host>\t Add host of <domain> to database.")

				os.Exit(0)
			}
			//handleHostCommand(os.Args[2:], dsn)
			hostCmd.Parse(os.Args[2:])
			fmt.Println("subcommand 'host'")
			fmt.Println("  dsn:", *hostDsn)
			fmt.Println("  tail:", hostCmd.Args())
			cmd.HandleHost(hostCmd, dsn)

			fmt.Println("")
		default:
			fmt.Println("dynpower-cli is a small helper tool to manage the dynpower database.")

			fmt.Println("\nUnknown or undefined command, please use the following commands:\n")
			fmt.Println("\tencrypt <password> :\t Encrypt password string to enter into database table.\n")

			fmt.Println("\tdomain [-dsn <dsn>] [command] [options]:\t Manage domain entries.")
			fmt.Println("\thost [-dsn <dsn>] [command] [options]:\t Manage host entries.")

			fmt.Println("\n\n")
			os.Exit(0)
			return


		}*/

}
