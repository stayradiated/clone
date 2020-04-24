package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var shallow = flag.Bool("shallow", false, "only fetch a single commit")
var https = flag.Bool("https", false, "use https instead of ssh")
var tag = flag.String("tag", "", "checkout a specific tag")

func git(args ...string) *exec.Cmd {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}

func main() {
	flag.Parse()

	args := flag.Args()

	var source string
	if len(args) < 1 {
		fmt.Println("Error: Must supply git repo.")
		os.Exit(1)
	} else {
		source = args[0]
	}

	var rootDir string
	if len(args) < 2 {
		rootDir = os.Getenv("HOME")
	} else {
		rootDir = args[1]
	}

	source = strings.TrimLeft(source, "https://")

	split := strings.Split(source, "/")
	if len(split) < 3 {
		fmt.Println("Error: Invalid git repo.")
		os.Exit(1)
	}

	host := split[0]
	user := split[1]
	repo := split[2]
	var url string

	url = "https://" + host + "/" + user + "/" + repo
	if (*https == false) {
		switch host {
		case "github.com":
			url = "git@github.com:" + user + "/" + repo
		case "bitbucket.org":
			url = "git@bitbucket.org:" + user + "/" + repo
		}
	}

	dir := fmt.Sprintf("%s/src/%s/%s/%s", rootDir, host, user, repo)

	exec.Command("mkdir", "-p", dir).Run()

	if (*shallow) {
		git("clone", "--depth", "1", url, dir).Run()
	} else {
		git("clone", url, dir).Run()
	}

	if (len(*tag) > 0) {
		fetchCmd := git("fetch", "--depth", "1", "origin", "tag", *tag)
		fetchCmd.Dir = dir
		fetchCmd.Run()

		resetCmd := git("reset", "--hard", *tag)
		resetCmd.Dir = dir
		resetCmd.Run()
	}
}
