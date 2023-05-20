package web

import (
	// "log"
	"net/http"
	"strconv"

	// "github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
	"github.com/aceberg/git-syr/internal/sync"
	"github.com/aceberg/git-syr/internal/yaml"
)

func reload() {

	close(AppConfig.Quit)

	yaml.Write(AppConfig.YamlPath, AllRepos)

	AppConfig.Quit = make(chan bool)
	sync.AllRepos(AllRepos, AppConfig.Quit)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var repo models.Repo

	repo.Name = r.FormValue("name")
	repo.Path = r.FormValue("path")
	repo.Pull = r.FormValue("pull")
	repo.Push = r.FormValue("push")
	repo.Timeout = r.FormValue("timeout")

	AllRepos = append(AllRepos, repo)

	// log.Println("AllRepos =", AllRepos)

	reload()

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var repo models.Repo

	idStr := r.FormValue("id")
	repo.ID, _ = strconv.Atoi(idStr)

	repo.Name = r.FormValue("name")
	repo.Path = r.FormValue("path")
	repo.Pull = r.FormValue("pull")
	repo.Push = r.FormValue("push")
	repo.Timeout = r.FormValue("timeout")
	repo.AddPush = r.FormValue("addpush")

	var newRepos []models.Repo
	for i := range AllRepos {
		if AllRepos[i].ID == repo.ID {
			if repo.Path != "" {
				newRepos = append(newRepos, repo)
			}
		} else {
			newRepos = append(newRepos, AllRepos[i])
		}
	}
	AllRepos = newRepos

	// log.Println("newRepos =", AllRepos)

	reload()

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
