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
	myRepo.Path = "/home/data/repo/00-public/drone"
	myRepo.Pull = "yes"

	git.Pull(myRepo.Path)
	if git.CheckIfPush(myRepo.Path) {
		git.Push(myRepo.Path)
	}
}

// func gitRepo(repo Repo) {
// 	for {
// 		fmt.Println(time.Now())
// 		time.Sleep(time.Duration(repo.Timeout) * time.Second) // Timeout
// 	}
// }
