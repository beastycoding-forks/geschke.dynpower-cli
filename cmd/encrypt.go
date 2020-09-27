package cmd

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

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
