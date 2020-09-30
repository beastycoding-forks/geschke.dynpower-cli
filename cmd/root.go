package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dynpower-cli",
	Short: "Manage dynpower database",
	Long: `
 dynpower-cli is a small helper tool to manage the dynpower database.
 `,
}

func Exec() {

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
