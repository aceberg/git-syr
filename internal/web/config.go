package web

import (
	"html/template"
	"net/http"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/conf"
	"github.com/aceberg/git-syr/internal/models"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig
	guiData.Icon = Icon

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "flatly", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "pulse", "quartz", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "yeti", "zephyr"}

	tmpl, err := template.ParseFiles("../../web/templates/config.html", "../../web/templates/header.html", "../../web/templates/footer.html")
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "config", guiData)
	check.IfError(err)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Theme = r.FormValue("theme")
	conf.Write(ConfigPath, AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
