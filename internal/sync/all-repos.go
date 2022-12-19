package sync

import (
	"log"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/git"
	"github.com/aceberg/git-syr/internal/models"
)

// AllRepos - sync all repos
func AllRepos(allRepos []models.Repo, quit chan bool) {
	var err error
	for _, oneRepo := range allRepos {
		oneRepo.Timeout, err = check.TimeToSec(oneRepo.Timeout)
		if err != nil {
			log.Println("ERROR:", oneRepo.Name, err)
			continue
		}
		if !git.CheckIfRepo(oneRepo.Path) {
			log.Println("ERROR:", oneRepo.Path, "is not a git repository")
			continue
		}
		if oneRepo.Push == "yes" || oneRepo.Pull == "yes" {
			log.Println("INFO: started sync for repo", oneRepo.Name)
			go syncRepo(oneRepo, quit)
		} else {
			log.Println("WARNING: repo", oneRepo.Name, "has no pull or push settings")
		}
	}
}
