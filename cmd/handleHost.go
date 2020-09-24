package cmd

import (
	"flag"
	"fmt"
)

/*
* Handle host-related commands
 */
func HandleHost(fs *flag.FlagSet, dsn string) {
	fmt.Println("handle Host command")
	fmt.Println(fs.Args())
	switch fs.Arg(0) {
	case "list":
		fmt.Println("Command: list")
	case "add":
		fmt.Println("Command: add")
	default:
		fmt.Println("Unknown command.")
	}

}
