package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
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

func browse(args []string) error {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return err
	}

	re := regexp.MustCompile("^git@([^:]+):(.+)\\.git")
	m := re.FindStringSubmatch(string(out))

	uri := fmt.Sprintf("https://%s/%s", m[1], m[2])
	return exec.Command("open", uri).Run()
}

func usage(args []string) error {
	fmt.Printf("usage: %s <command>\n", args[0])
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("    browse    open repository in browser")
	os.Exit(1)
	return nil
}
