package git

import (
	. "github.com/aceberg/GitSyncTimer/internal/common"
	"log"
	"os/exec"
	"time"
)

func Push(path string) {

	log.Println("Push repo", path)

	currentTime := time.Now()
	timeString := string(currentTime.Format("2006-01-02 15:04:05"))

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "add", "-A")
	out, err := cmd.CombinedOutput()
	CheckIfError(err)
	log.Println(string(out))

	cmd = exec.Command("git", gitDir, gitTree, "commit", "-m", timeString)
	out, err = cmd.CombinedOutput()
	CheckIfError(err)
	log.Println(string(out))

	cmd = exec.Command("git", gitDir, gitTree, "push")
	out, err = cmd.CombinedOutput()
	CheckIfError(err)
	log.Println(string(out))
}
