package main

import (
	"fmt"
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
)

//. "github.com/go-git/go-git/v5/_examples"
func main() {

	//fmt.Printf("\x1b[31;1m%s\x1b[0m\n", "Welcome to color text")
	//CheckArgs("<url>", "<directory>", "<github_username>", "<github_password>")

	//url, directory, username, password := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	var url string = "https://github.com/mateors/email"
	currDir, _ := os.Getwd()
	repoDir := filepath.Join(currDir, "email")
	var directory string = repoDir
	var username string = "mateors"
	var password string = ""

	// Clone the given repository to the given directory
	Info("git clone %s %s", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: username,
			Password: password,
		},
		URL:           url,
		ReferenceName: plumbing.ReferenceName("refs/heads/main"),
		Progress:      os.Stdout,
	})
	if err != nil {
		fmt.Println(">", err)
		os.Exit(1)
	}
	//CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
