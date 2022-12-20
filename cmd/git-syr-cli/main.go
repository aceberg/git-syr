package main

import (
	"flag"
	"time"
	// "log"

	"github.com/aceberg/git-syr/internal/cli"
	"github.com/aceberg/git-syr/internal/logfile"
)

const yamlPath = "/data/git-syr/repos.yaml"
const logPath = ""

func main() {
	logPtr := flag.String("l", logPath, "Path to log file")
	yamlPtr := flag.String("r", yamlPath, "Path to repos yaml file")
	flag.Parse()

	go logfile.Output(*logPtr)
	time.Sleep(1 * time.Second)

	cli.Start(*yamlPtr)
}
