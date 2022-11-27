package main

import (
	"log"
	"strconv"
	"time"
	// . "github.com/aceberg/GitSyncTimer/internal/common"
	"github.com/aceberg/GitSyncTimer/internal/git"
	. "github.com/aceberg/GitSyncTimer/internal/models"
	"github.com/aceberg/GitSyncTimer/internal/yaml"
)

const yamlPath = "/data/GitSyncTimer/repos.yaml"

func main() {
	allRepos := yaml.ReadYaml(yamlPath)

	log.Println("INFO: all repos", allRepos)

	for _, oneRepo := range allRepos {
		log.Println("INFO: started sync for repo", oneRepo.Name)
		go gitRepo(oneRepo)
	}

	select {}
}

func gitRepo(repo Repo) {
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
