package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	var source string
	if len(os.Args) < 2 {
		fmt.Println("Error: Must supply git repo.")
		os.Exit(1)
	} else {
		source = os.Args[1]
	}

	var rootDir string
	if len(os.Args) < 3 {
		rootDir = os.Getenv("HOME")
	} else {
		rootDir = os.Args[2]
	}

	split := strings.Split(source, "/")

	if len(split) < 3 {
		fmt.Println("Error: Invalid git repo.")
		os.Exit(1)
	}

	host := split[0]
	user := split[1]
	repo := split[2]
	var url string

	switch host {
	case "github.com":
		url = "git@github.com:" + user + "/" + repo
	case "bitbucket.org":
		url = "git@bitbucket.org:" + user + "/" + repo
	default:
		url = "https://" + host + "/" + user + "/" + repo
	}

	dir := fmt.Sprintf("%s/src/%s/%s/%s", rootDir, host, user, repo)

	exec.Command("mkdir", "-p", dir).Run()
	exec.Command("fasd", "-A", dir).Run()

	cmd := exec.Command("git", "clone", url, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
