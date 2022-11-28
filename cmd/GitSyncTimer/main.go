package main

import (
	"flag"
	"github.com/aceberg/GitSyncTimer/internal/check"
	"github.com/aceberg/GitSyncTimer/internal/git"
	"github.com/aceberg/GitSyncTimer/internal/sync"
	"github.com/aceberg/GitSyncTimer/internal/yaml"
	"log"
)

// const yamlPath = "/data/GitSyncTimer/repos.yaml"

func main() {
	yamlPtr := flag.String("r", "/data/GitSyncTimer/repos.yaml", "Path to repos yaml file")
	webPtr := flag.Bool("w", false, "Launch web gui")
	flag.Parse()
	yamlPath := *yamlPtr

	allRepos := yaml.ReadYaml(yamlPath)
	log.Println("INFO: all repos", allRepos)

	if allRepos == nil {
		log.Printf("ERROR: no repos")
		return
	}

	var err error
	for _, oneRepo := range allRepos {
		oneRepo.Data.Timeout, err = check.Timeout(oneRepo.Data.Timeout)
		if err != nil {
			log.Println("ERROR:", oneRepo.Name, err)
		} else {
			if git.CheckIfRepo(oneRepo.Data.Path) {
				log.Println("INFO: started sync for repo", oneRepo.Name)
				go sync.SyncRepo(oneRepo)
			} else {
				log.Println("ERROR:", oneRepo.Data.Path, "is not a git repository")
			}
		}
	}

	if *webPtr {
		log.Println("WEB GUI")
	}

	select {}
}
