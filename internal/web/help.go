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

	guiData.Themes = []string{"● To delete repository save it with empty path",
		"● Timeout examples: '30s' - 30 seconds, '5m' - 5 minutes, '3h' - 3 hours"}

	tmpl, err := template.ParseFS(TemplHTML, "templates/log.html", "templates/header.html", "templates/footer.html")
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "log", guiData)
	check.IfError(err)
}
