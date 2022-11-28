package git

import (
	. "github.com/aceberg/GitSyncTimer/internal/common"
	"log"
	"os/exec"
)

// Pull - executes `git pull` in repo
func Pull(path string) {

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "pull")

	out, err := cmd.CombinedOutput()
	CheckIfError(err)

	log.Println("INFO: Pull repo", path, "\n", string(out))
}
