package git

import (
	"log"
	"os/exec"

	"github.com/aceberg/GitSyncTimer/internal/check"
)

// Pull - executes `git pull` in repo
func Pull(path string) {

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "pull")

	out, err := cmd.CombinedOutput()
	check.IfError(err)

	log.Println("INFO: Pull repo", path, "\n", string(out))
}
