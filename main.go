package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var https = flag.Bool("https", false, "use https instead of ssh")
var shallow = flag.Bool("shallow", false, "only fetch the most recent commit")
var tag = flag.String("tag", "", "checkout a specific tag")
var ref = flag.String("ref", "", "checkout a specific ref")

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
	if *https == false {
		switch host {
		case "github.com":
			url = "git@github.com:" + user + "/" + repo
		case "bitbucket.org":
			url = "git@bitbucket.org:" + user + "/" + repo
		}
	}

	dir := fmt.Sprintf("%s/src/%s/%s/%s", rootDir, host, user, repo)

	err := exec.Command("mkdir", "-p", dir).Run()
	if err != nil {
		log.Fatalf("Could not mkdir -p %s", dir)
	}

	hasRef := len(*ref) > 0
	hasTag := len(*tag) > 0

	if *shallow || hasRef || hasTag {
		err = git("clone", "--depth", "1", url, dir).Run()
	} else {
		err = git("clone", url, dir).Run()
	}
	if err != nil {
		log.Fatalf("Could not git clone %s", url)
	}

	if hasRef || hasTag {
		var fetchCmd *exec.Cmd
		if hasTag {
			fetchCmd = git("fetch", "--depth", "1", "origin", "tag", *tag)
		} else {
			fetchCmd = git("fetch", "--depth", "1", "origin", *ref)
		}

		fetchCmd.Dir = dir
		err := fetchCmd.Run()
		if err != nil {
			log.Fatal("Could not git fetch")
		}

		var resetCmd *exec.Cmd
		if hasTag {
			resetCmd = git("reset", "--hard", *tag)
		} else {
			resetCmd = git("reset", "--hard", *ref)
		}
		resetCmd.Dir = dir
		err = resetCmd.Run()
		if err != nil {
			log.Fatal("Could not git reset")
		}
	}
}
