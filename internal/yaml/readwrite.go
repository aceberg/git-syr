package yaml

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

// Read - read .yaml file to []struct
func Read(path string) []models.Repo {

	file, err := os.ReadFile(path)
	check.IfError(err)

	var allRepos []models.Repo
	err = yaml.Unmarshal(file, &allRepos)
	check.IfError(err)

	return allRepos
}

// Write - write []struct to  .yaml file
func Write(path string, allRepos []models.Repo) {

	yamlData, err := yaml.Marshal(&allRepos)
	check.IfError(err)

	err = os.WriteFile(path, yamlData, 0644)
	check.IfError(err)

	log.Println("INFO: writing new repos file to", path, "\n", string(yamlData))
}
