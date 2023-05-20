package git

import (
	"log"
	"os/exec"
	"time"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

// Push - executes
// git add -A
// git commit -m `date`
// git push
// in repo
func Push(repo models.Repo) {

	currentTime := time.Now()
	timeString := string(currentTime.Format("2006-01-02 15:04:05"))

	gitDir := "--git-dir=" + repo.Path + "/.git"
	gitTree := "--work-tree=" + repo.Path

	cmd := exec.Command("git", gitDir, gitTree, "add", "-A")
	out1, err := cmd.CombinedOutput()
	check.IfError(err)

	cmd = exec.Command("git", gitDir, gitTree, "commit", "-m", timeString)
	out2, err := cmd.CombinedOutput()
	check.IfError(err)

	if repo.AddPush != "" {
		cmd = exec.Command("git", gitDir, gitTree, "push", repo.AddPush)
	} else {
		cmd = exec.Command("git", gitDir, gitTree, "push")
	}
	out3, err := cmd.CombinedOutput()
	check.IfError(err)

	log.Println("INFO: Push repo", repo.Path, "\n", string(out1), "\n", string(out2), "\n", string(out3))
}
