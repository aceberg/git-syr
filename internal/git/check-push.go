package git

import (
	"log"
	"os/exec"

	"github.com/aceberg/git-syr/internal/check"
)

// CheckIfPush checks if the repo needs push
func CheckIfPush(path string) bool {

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "status", "-s")

	out, err := cmd.CombinedOutput()
	check.IfError(err)

	log.Println("INFO: Check if repo needs push", path, "\n", string(out))

	return string(out) != ""
}
