package sync

import (
	"github.com/aceberg/GitSyncTimer/internal/check"
	"github.com/aceberg/GitSyncTimer/internal/git"
	. "github.com/aceberg/GitSyncTimer/internal/models"
	"log"
)

func AllRepos(allRepos []Repo) {
	var err error
	for _, oneRepo := range allRepos {
		oneRepo.Data.Timeout, err = check.Timeout(oneRepo.Data.Timeout)
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
