package sync

import (
	// "log"
	"strconv"
	"time"

	"github.com/aceberg/git-syr/internal/git"
	"github.com/aceberg/git-syr/internal/models"
)

func syncRepo(repo models.Repo, quit chan bool) {
	var lastDate time.Time

	timeout, _ := strconv.Atoi(repo.Timeout)

	for {
		select {
		case <-quit:
			return
		default:
			nowDate := time.Now()
			plusDate := lastDate.Add(time.Duration(timeout) * time.Second)

			if nowDate.After(plusDate) {
				if repo.Pull == "yes" {
					git.Pull(repo.Path)
				}
				if repo.Push == "yes" && git.CheckIfPush(repo.Path) {
					git.Push(repo.Path)
				}
				lastDate = time.Now()
			}

			time.Sleep(time.Duration(1) * time.Second) // Cycle Timeout to check if quit
		}
	}
}
