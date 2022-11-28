package sync

import (
	"github.com/aceberg/GitSyncTimer/internal/git"
	. "github.com/aceberg/GitSyncTimer/internal/models"
	"log"
	"strconv"
	"time"
)

func SyncRepo(repo Repo) {
	for {
		check := true
		if repo.Data.Pull == "yes" {
			check = false
			git.Pull(repo.Data.Path)
		}
		if repo.Data.Push == "yes" && git.CheckIfPush(repo.Data.Path) {
			check = false
			git.Push(repo.Data.Path)
		}
		if check {
			log.Println("WARNING: repo", repo.Name, "has no pull or push settings")
			break
		}

		timeout, _ := strconv.Atoi(repo.Data.Timeout)
		time.Sleep(time.Duration(timeout) * time.Second) // Timeout
	}
}
