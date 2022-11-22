package git

import (
	. "github.com/aceberg/GitSyncTimer/internal/common"
	"log"
	"os/exec"
)

func Pull(path string) {

	log.Println("Pull repo", path)

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "pull")

	out, err := cmd.CombinedOutput()
	CheckIfError(err)

	log.Println(string(out))
}
