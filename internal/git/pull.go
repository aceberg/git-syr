package git

import (
	"fmt"
	. "github.com/aceberg/GitBackup/internal/common"
	"github.com/go-git/go-git/v5"
)

func Pull(path string) {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(path)
	CheckIfError(err)

	// Get the working directory for the repository
	w, err := r.Worktree()
	CheckIfError(err)

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	CheckIfError(err)

	fmt.Println("PULL")
}
