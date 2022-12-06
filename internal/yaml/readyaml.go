package yaml

import (
	// "fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/aceberg/GitSyncTimer/internal/check"
	"github.com/aceberg/GitSyncTimer/internal/models"
)

// ReadYaml - read .yaml file to []struct
func ReadYaml(path string) []models.Repo {

	file, err := os.ReadFile(path)
	check.IfError(err)

	var allRepos []models.Repo
	err = yaml.Unmarshal(file, &allRepos)
	check.IfError(err)

	return allRepos
}
