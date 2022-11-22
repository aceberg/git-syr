package main

import (
	// "fmt"
	// "time"
	// . "github.com/aceberg/GitSyncTimer/internal/common"
	"github.com/aceberg/GitSyncTimer/internal/git"
	. "github.com/aceberg/GitSyncTimer/internal/models"
)

func main() {
	var myRepo Repo
	myRepo.Timeout = "120"
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
