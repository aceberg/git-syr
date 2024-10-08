package main

import (
	"flag"
	"time"
	// "log"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/cli"
	"github.com/aceberg/git-syr/internal/logfile"
	"github.com/aceberg/git-syr/internal/web"
)

const confPath = "config.yaml"
const yamlPath = "repos.yaml"
const logPath = "git-syr.log"
const nodePath = ""

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	logPtr := flag.String("l", logPath, "Path to log file")
	nodePtr := flag.String("n", nodePath, "Path to local node modules")
	yamlPtr := flag.String("r", yamlPath, "Path to repos yaml file")
	webPtr := flag.Bool("w", false, "Launch without web gui")
	flag.Parse()

	check.Path(*confPtr)
	check.Path(*logPtr)
	check.Path(*yamlPtr)

	go logfile.Output(*logPtr)
	time.Sleep(1 * time.Second)

	if *webPtr {
		cli.Start(*yamlPtr)
	} else {
		web.Gui(*confPtr, *yamlPtr, *logPtr, *nodePtr)
	}
}
