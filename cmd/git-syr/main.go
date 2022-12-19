package main

import (
	"flag"
	"time"
	// "log"

	"github.com/aceberg/git-syr/internal/cli"
	"github.com/aceberg/git-syr/internal/logfile"
	"github.com/aceberg/git-syr/internal/web"
)

const confPath = "/data/git-syr/config.yaml"
const yamlPath = "/data/git-syr/repos.yaml"
const logPath = "/data/git-syr/git-syr.log"

func main() {
	confPtr := flag.String("c", confPath, "Path to config yaml file")
	logPtr := flag.String("l", logPath, "Path to log file")
	yamlPtr := flag.String("r", yamlPath, "Path to repos yaml file")
	webPtr := flag.Bool("w", true, "Launch web gui")
	flag.Parse()

	go logfile.Output(*logPtr)
	time.Sleep(1 * time.Second)

	if *webPtr {
		web.Gui(*confPtr, *yamlPtr, *logPtr)
	} else {
		cli.Start(*yamlPtr, *logPtr)
	}
}
