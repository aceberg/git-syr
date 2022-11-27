package git

import (
	. "github.com/aceberg/GitSyncTimer/internal/common"
	"log"
	"os/exec"
)

// CheckIfPush checks if the repo needs push
func CheckIfPush(path string) bool {

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "status", "-s")

	out, err := cmd.CombinedOutput()
	CheckIfError(err)

	log.Println("INFO: Check if repo needs push", path, "\n", string(out))

	if string(out) == "" {
		return false
	}
	
	return true
}
