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

	if allRepos == nil {
		log.Printf("ERROR: no repos")
		return
	}

	for _, oneRepo := range allRepos {
		if git.CheckIfRepo(oneRepo.Data.Path) {
			log.Println("INFO: started sync for repo", oneRepo.Name)
			go gitRepo(oneRepo)
		} else {
			log.Println("ERROR:", oneRepo.Data.Path, "is not a git repository")
		}
	}

	select {}
}

func gitRepo(repo Repo) {
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
			break
		}

		timeout, _ := strconv.Atoi(repo.Data.Timeout)
		time.Sleep(time.Duration(timeout) * time.Second) // Timeout
	}
}
