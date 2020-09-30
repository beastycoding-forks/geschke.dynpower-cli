package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var hostDsn string

func init() {
	hostCmd.Flags().StringVarP(&hostDsn, "dsn", "d", "", "MySQL/MariaDB Data Source Name as described in https://github.com/go-sql-driver/mysql")

	//domainCmd.AddCommand(domainListCmd)
	rootCmd.AddCommand(hostCmd)

}

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Handle host...Show details for an article",
	Long:  `Details for an article`,
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command,
		args []string) {
		handleHost()
	},
}

/*
* Handle host-related commands
 */
func handleHost() {
	fmt.Println("handle Host command")
	/*fmt.Println(fs.Args())
	switch fs.Arg(0) {
	case "list":
		fmt.Println("Command: list")
	case "add":
		fmt.Println("Command: add")
	default:
		fmt.Println("Unknown command.")
	}*/

}
