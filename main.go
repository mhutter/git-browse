package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return
	}

	cmd, err := openCmd()
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	uri, err := extractURL(strings.TrimSpace(string(out)))
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	exec.Command(cmd, uri).Start()
}

func openCmd() (string, error) {
	candidates := []string{"open", "xdg-open"}
	for _, cmd := range candidates {
		if cmdExists(cmd) {
			return cmd, nil
		}
	}

	return "", fmt.Errorf("I don't know how to open a browser! tried: %s",
		candidates)
}

func cmdExists(exe string) bool {
	return exec.Command("which", exe).Run() == nil
}

func extractURL(remote string) (string, error) {
	patterns := []string{
		"^git@([^:]+):((?U).+)(?:\\.git)?$",
		"^https://([^/]+)/((?U).+)(?:\\.git)?$",
	}
	for _, p := range patterns {
		re := regexp.MustCompile(p)
		m := re.FindStringSubmatch(remote)
		if len(m) == 3 {
			return fmt.Sprintf("https://%s/%s", m[1], m[2]), nil
		}
	}

	return "", fmt.Errorf("Could not parse remote '%s'", remote)
}
