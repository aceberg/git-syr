package main

import (
	"fmt"
	// "time"
	// . "github.com/aceberg/GitSyncTimer/internal/common"
	// "github.com/aceberg/GitSyncTimer/internal/git"
	// . "github.com/aceberg/GitSyncTimer/internal/models"
	"github.com/aceberg/GitSyncTimer/internal/yaml"
)

const configPath = "/data/GitSyncTimer/"
const yamlPath = "/data/GitSyncTimer/repos.yaml"

func main() {
	allRepos := yaml.ReadYaml(yamlPath)

	fmt.Println("REPO:", allRepos)

	// git.Pull(myRepo.Path)
	// if git.CheckIfPush(myRepo.Path) {
	// 	git.Push(myRepo.Path)
	// }
}

// func gitRepo(repo Repo) {
// 	for {
// 		fmt.Println(time.Now())
// 		time.Sleep(time.Duration(repo.Timeout) * time.Second) // Timeout
// 	}
// }
