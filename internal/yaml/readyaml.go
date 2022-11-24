package yaml

import (
	// "fmt"
	. "github.com/aceberg/GitSyncTimer/internal/common"
	. "github.com/aceberg/GitSyncTimer/internal/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadYaml(path string) []Repo {

	file, err := ioutil.ReadFile(path)
	CheckIfError(err)

	var allRepos []Repo
	err = yaml.Unmarshal(file, &allRepos)
	CheckIfError(err)

	return allRepos
}
