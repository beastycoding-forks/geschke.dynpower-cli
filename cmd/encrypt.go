package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt [access key]",
	Short: "Encrypt access key string to enter into database table.",
	//Long:  `All software has versions. This is Hugo's`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command,
		args []string) {
		Encrypt(args[0])

	},
}

// HashPassword hashes a submitted string with bcrypt algorithm to store as password in database
func HashPassword(password string) (string, error) {
	passWordBytes := []byte(password)
	bytes, err := bcrypt.GenerateFromPassword(passWordBytes, bcrypt.DefaultCost)
	return string(bytes), err
}

// Encrypt string with bcrypt algorithm
func Encrypt(pw string) {
	fmt.Println("Password: ")
	fmt.Println(pw)

	hashedPassword, err := HashPassword(pw)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("\nYour encrypted password:\n")
	fmt.Println(hashedPassword)
	fmt.Println("\nPlease enter this string in the field 'access_key' into the domains table.\n")

	// test checking:
	//err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	//fmt.Println(err) // nil means it is a match

}
