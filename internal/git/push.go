package git

import (
	"log"
	"os/exec"
	"time"

	"github.com/aceberg/GitSyncTimer/internal/check"
)

// Push - executes
// git add -A
// git commit -m `date`
// git push
// in repo
func Push(path string) {

	currentTime := time.Now()
	timeString := string(currentTime.Format("2006-01-02 15:04:05"))

	gitDir := "--git-dir=" + path + "/.git"
	gitTree := "--work-tree=" + path

	cmd := exec.Command("git", gitDir, gitTree, "add", "-A")
	out1, err := cmd.CombinedOutput()
	check.IfError(err)

	cmd = exec.Command("git", gitDir, gitTree, "commit", "-m", timeString)
	out2, err := cmd.CombinedOutput()
	check.IfError(err)

	cmd = exec.Command("git", gitDir, gitTree, "push")
	out3, err := cmd.CombinedOutput()
	check.IfError(err)

	log.Println("INFO: Push repo", path, "\n", string(out1), "\n", string(out2), "\n", string(out3))
}
