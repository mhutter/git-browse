package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var fn func([]string) error

	cmd := "usage"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	switch cmd {
	case "browse":
		fn = browse
	default:
		fn = usage
	}

	if err := fn(os.Args); err != nil {
		log.Printf("ERROR: %s", err)
	}
}

func usage(args []string) error {
	fmt.Printf("usage: %s <command>\n", args[0])
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("    browse    open repository in browser")
	os.Exit(1)
	return nil
}
