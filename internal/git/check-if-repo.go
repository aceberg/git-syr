package git

import (
	"os"
)

// CheckIfRepo checks if the path is git repo
func CheckIfRepo(path string) bool {

	_, err := os.Stat(path + "/.git")

	return err == nil
}
