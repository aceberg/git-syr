package main

import (
	"flag"
	"log"

	"github.com/aceberg/git-syr/internal/sync"
	"github.com/aceberg/git-syr/internal/web"
	"github.com/aceberg/git-syr/internal/yaml"
)

const confPath = "/data/git-syr/config.yaml"
const yamlPath = "/data/git-syr/repos.yaml"

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	yamlPtr := flag.String("r", yamlPath, "Path to repos yaml file")
	webPtr := flag.Bool("w", true, "Launch web gui")
	flag.Parse()

	allRepos := yaml.Read(*yamlPtr)
	log.Println("INFO: all repos", allRepos)

	if allRepos == nil {
		log.Printf("ERROR: no repos")
		return
	}

	sync.AllRepos(allRepos)

	if *webPtr {
		web.Gui(*confPtr, *yamlPtr, allRepos)
	}

	select {}
}
