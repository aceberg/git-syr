package web

import (
	"net/http"

	"github.com/aceberg/git-syr/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	for i := range AllRepos {
		AllRepos[i].ID = i + 1
	}

	guiData.Repos = AllRepos

	execTemplate(w, "index", guiData)
}
