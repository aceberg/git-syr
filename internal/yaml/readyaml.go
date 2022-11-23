package yaml

import (
	// "fmt"
	. "github.com/aceberg/GitSyncTimer/internal/common"
	. "github.com/aceberg/GitSyncTimer/internal/models"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadYaml(path string) []OneRepo {

	yfile, err := ioutil.ReadFile(path)
	CheckIfError(err)

	var allRepos []OneRepo
	err = yaml.Unmarshal(yfile, &allRepos)
	CheckIfError(err)

	return allRepos
}
