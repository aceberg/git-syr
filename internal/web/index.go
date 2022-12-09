package web

import (
	"html/template"
	"net/http"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon
	guiData.Repos = AllRepos

	tmpl, err := template.ParseFiles("../../web/templates/index.html", "../../web/templates/header.html", "../../web/templates/footer.html")
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "index", guiData)
	check.IfError(err)
}
