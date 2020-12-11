package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	help = `
		data-tree-service has all endpoints related to  data-tree information.
		
		The following sub-commands are supported
		  1. serve - Run the application as a server.
		  3. help - Prints this help.
		  4. version - Prints the version of the application.
		
		You can pass the following configuration through environment variables.
		The configuration, unless stated otherwise, are mandatory.
	`
)

var version = "v1"

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("No arguments provided.")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch strings.ToLower(cmd) {
	case "serve":
		serve()
	case "help":
		fmt.Print(help)
	case "version":
		fmt.Println(version)
	default:
		fmt.Printf("Unknown command %q", cmd)
		os.Exit(1)
	}
}

func serve() {
	fmt.Println("hello")

}