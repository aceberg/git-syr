package main

import (
	"flag"
	"github.com/aceberg/GitSyncTimer/internal/sync"
	"github.com/aceberg/GitSyncTimer/internal/web"
	"github.com/aceberg/GitSyncTimer/internal/yaml"
	"log"
)

const confPath = "/data/GitSyncTimer/config.yaml"
const yamlPath = "/data/GitSyncTimer/repos.yaml"

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	yamlPtr := flag.String("r", yamlPath, "Path to repos yaml file")
	webPtr := flag.Bool("w", true, "Launch web gui")
	flag.Parse()

	allRepos := yaml.ReadYaml(*yamlPtr)
	log.Println("INFO: all repos", allRepos)

	if allRepos == nil {
		log.Printf("ERROR: no repos")
		return
	}

	sync.AllRepos(allRepos)

	if *webPtr {
		web.Gui(*confPtr, allRepos)
	}

	select {}
}
