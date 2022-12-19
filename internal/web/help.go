package web

import (
	"html/template"
	"net/http"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

func helpHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.Themes = []string{"● cerulean", "● cosmo"}

	tmpl, err := template.ParseFiles("../../web/templates/log.html", "../../web/templates/header.html", "../../web/templates/footer.html")
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "log", guiData)
	check.IfError(err)
}
