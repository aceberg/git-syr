package sync

import (
	// "log"
	"strconv"
	"time"

	"github.com/aceberg/GitSyncTimer/internal/git"
	"github.com/aceberg/GitSyncTimer/internal/models"
)

func syncRepo(repo models.Repo) {
	for {
		if repo.Data.Pull == "yes" {
			git.Pull(repo.Data.Path)
		}
		if repo.Data.Push == "yes" && git.CheckIfPush(repo.Data.Path) {
			git.Push(repo.Data.Path)
		}

		timeout, _ := strconv.Atoi(repo.Data.Timeout)
		time.Sleep(time.Duration(timeout) * time.Second) // Timeout
	}
}
