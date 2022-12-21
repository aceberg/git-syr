package main

import (
	"flag"
	"time"
	// "log"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/cli"
	"github.com/aceberg/git-syr/internal/logfile"
)

const yamlPath = "repos.yaml"
const logPath = ""

func main() {
	logPtr := flag.String("l", logPath, "Path to log file")
	yamlPtr := flag.String("r", yamlPath, "Path to repos yaml file")
	flag.Parse()

	check.Path(*logPtr)
	check.Path(*yamlPtr)

	go logfile.Output(*logPtr)
	time.Sleep(1 * time.Second)

	cli.Start(*yamlPtr)
}
