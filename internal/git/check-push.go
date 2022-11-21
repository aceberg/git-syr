package git

import (
	. "github.com/aceberg/GitBackup/internal/common"
	"log"
	"os/exec"
)

func CheckIfPush(path string) bool {

	log.Println("Check if repo needs push", path)

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "status", "-s")

	out, err := cmd.CombinedOutput()
	CheckIfError(err)

	log.Println(string(out))

	if string(out) == "" {
		return false
	} else {
		return true
	}
}
