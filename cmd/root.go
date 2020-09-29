package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Manage domain entries",
	Long:  `List first 5 news`,
	Run: func(cmd *cobra.Command,
		args []string) {
		HandleDomain()
	},
}

var cmdDomainList = &cobra.Command{
	Use:   "list",
	Short: "List domains in database",
	Long:  `List all domains in the dynpower database`,
	Run: func(cmd *cobra.Command,
		args []string) {
		List(args[0])
	},
}

var cmdHost = &cobra.Command{
	Use:   "host [id]",
	Short: "Handle host...Show details for an article",
	Long:  `Details for an article`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command,
		args []string) {
		HandleHost(args[0])
	},
}

var rootCmd = &cobra.Command{
	Use:   "dynpower-cli",
	Short: "Manage dynpower database",
	Long: `
 dynpower-cli is a small helper tool to manage the dynpower database.
 `,
}

func Exec() {
	domainCmd.AddCommand(cmdDomainList)
	var Dsn string
	domainCmd.Flags().StringVarP(&Dsn, "dsn", "d", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql")
	rootCmd.AddCommand(domainCmd)
	rootCmd.AddCommand(cmdHost)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
