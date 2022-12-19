package sync

import (
	// "log"
	"strconv"
	"time"

	"github.com/aceberg/git-syr/internal/git"
	"github.com/aceberg/git-syr/internal/models"
)

func syncRepo(repo models.Repo, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			if repo.Pull == "yes" {
				git.Pull(repo.Path)
			}
			if repo.Push == "yes" && git.CheckIfPush(repo.Path) {
				git.Push(repo.Path)
			}

			timeout, _ := strconv.Atoi(repo.Timeout)
			time.Sleep(time.Duration(timeout) * time.Second) // Timeout
		}
	}
}
