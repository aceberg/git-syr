package cli

import (
	"log"

	"github.com/aceberg/git-syr/internal/sync"
	"github.com/aceberg/git-syr/internal/yaml"
)

// Start cli app
func Start(yamlPath string) {
	log.Println("INFO: starting cli app with", yamlPath)

	allRepos := yaml.Read(yamlPath)
	log.Println("INFO: all repos", allRepos)

	if allRepos == nil {
		log.Printf("ERROR: no repos")
		return
	}

	quit := make(chan bool)
	sync.AllRepos(allRepos, quit)

	select {}
}
