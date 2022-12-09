package sync

import (
	"log"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/git"
	"github.com/aceberg/git-syr/internal/models"
)

// AllRepos - sync all repos
func AllRepos(allRepos []models.Repo) {
	var err error
	for _, oneRepo := range allRepos {
		oneRepo.Data.Timeout, err = check.TimeToSec(oneRepo.Data.Timeout)
		if err != nil {
			log.Println("ERROR:", oneRepo.Name, err)
			continue
		}
		if !git.CheckIfRepo(oneRepo.Data.Path) {
			log.Println("ERROR:", oneRepo.Data.Path, "is not a git repository")
			continue
		}
		if oneRepo.Data.Push == "yes" || oneRepo.Data.Pull == "yes" {
			log.Println("INFO: started sync for repo", oneRepo.Name)
			go syncRepo(oneRepo)
		} else {
			log.Println("WARNING: repo", oneRepo.Name, "has no pull or push settings")
		}
	}
}
