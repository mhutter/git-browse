package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

func browse(args []string) error {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return err
	}

	re := regexp.MustCompile("^git@([^:]+):(.+)\\.git")
	m := re.FindStringSubmatch(string(out))

	uri := fmt.Sprintf("https://%s/%s", m[1], m[2])
	return exec.Command(openCmd(), uri).Run()
}

func openCmd() string {
	candidates := []string{"open", "xdg-open"}
	for _, cmd := range candidates {
		if cmdExists(cmd) {
			return cmd
		}
	}
	log.Fatalf("I don't know how to open a browser! tried: %s", candidates)
	return ""
}

func cmdExists(exe string) bool {
	return exec.Command("which", exe).Run() == nil
}
