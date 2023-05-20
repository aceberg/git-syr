package web

import (
	// "log"
	"net/http"
	"os"
	"strconv"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	guiData.Config = AppConfig
	guiData.Themes = make([]string, 2)

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	check.IfError(err)

	for _, oneRepo := range AllRepos {
		if oneRepo.ID == id {
			guiData.Path = oneRepo.Path

			file1, err := os.ReadFile(oneRepo.Path + "/.git/config")
			check.IfError(err)
			guiData.Themes[0] = string(file1)

			file2, err := os.ReadFile(oneRepo.Path + "/.gitignore")
			check.IfError(err)
			guiData.Themes[1] = string(file2)
		}
	}

	execTemplate(w, "edit", guiData)
}

func saveFileHandler(w http.ResponseWriter, r *http.Request) {

	file := r.FormValue("file")
	path := r.FormValue("path")
	text := r.FormValue("text")

	if file == "config" {
		err := os.WriteFile(path+"/.git/config", []byte(text), 0644)
		check.IfError(err)
	}
	if file == "ignore" {
		err := os.WriteFile(path+"/.gitignore", []byte(text), 0644)
		check.IfError(err)
	}

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
