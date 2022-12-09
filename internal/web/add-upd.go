package web

import (
	"log"
	"net/http"
	"strconv"

	// "github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
	"github.com/aceberg/git-syr/internal/yaml"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	var repo models.Repo

	repo.Name = r.FormValue("name")
	repo.Data.Path = r.FormValue("path")
	repo.Data.Pull = r.FormValue("pull")
	repo.Data.Push = r.FormValue("push")
	repo.Data.Timeout = r.FormValue("timeout")

	AllRepos = append(AllRepos, repo)

	// log.Println("AllRepos =", AllRepos)

	yaml.Write(YamlPath, AllRepos)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var repo models.Repo

	idStr := r.FormValue("id")
	repo.ID, _ = strconv.Atoi(idStr)

	repo.Name = r.FormValue("name")
	repo.Data.Path = r.FormValue("path")
	repo.Data.Pull = r.FormValue("pull")
	repo.Data.Push = r.FormValue("push")
	repo.Data.Timeout = r.FormValue("timeout")

	var newRepos []models.Repo
	for i, _ := range AllRepos {
		if AllRepos[i].ID == repo.ID {
			if repo.Data.Path != "" {
				newRepos = append(newRepos, repo)
			}
		} else {
			newRepos = append(newRepos, AllRepos[i])
		}
	}
	AllRepos = newRepos

	log.Println("newRepos =", AllRepos)

	yaml.Write(YamlPath, AllRepos)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
