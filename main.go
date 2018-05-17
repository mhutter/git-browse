package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return
	}

	re := regexp.MustCompile("^git@([^:]+):(.+)\\.git")
	m := re.FindStringSubmatch(string(out))

	uri := fmt.Sprintf("https://%s/%s", m[1], m[2])
	cmd, err := openCmd()
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
