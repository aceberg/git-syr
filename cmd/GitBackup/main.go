package main

import (
	// "fmt"
	// "time"
	// . "github.com/aceberg/GitBackup/internal/common"
	"github.com/aceberg/GitBackup/internal/git"
)

type Repo struct {
	Path    string
	Timeout int
	Pull    string
	Push    string
}

func main() {
	var myRepo Repo
	myRepo.Timeout = 120
	myRepo.Path = "/home/data/repo/00-public/MyDocs"
	myRepo.Pull = "yes"

	git.Pull(myRepo.Path)
}

// func gitRepo(repo Repo) {
// 	for {
// 		// We instantiate a new repository targeting the given path (the .git folder)
// 		r, err := git.PlainOpen(repo.Path)
// 		CheckIfError(err)

// 		// Get the working directory for the repository
// 		w, err := r.Worktree()
// 		CheckIfError(err)

// 		// Pull the latest changes from the origin remote and merge into the current branch
// 		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
// 		CheckIfError(err)

// 		fmt.Println(time.Now())
// 		time.Sleep(time.Duration(repo.Timeout) * time.Second) // Timeout
// 	}
// }
