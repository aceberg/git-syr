package web

import (
	"net/http"

	"github.com/aceberg/git-syr/internal/conf"
	"github.com/aceberg/git-syr/internal/models"
)

func configHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	guiData.Themes = []string{"cerulean", "cosmo", "cyborg", "darkly", "emerald", "flatly", "grass", "grayscale", "journal", "litera", "lumen", "lux", "materia", "minty", "morph", "ocean", "pulse", "quartz", "sand", "sandstone", "simplex", "sketchy", "slate", "solar", "spacelab", "superhero", "united", "vapor", "wood", "yeti", "zephyr"}

	execTemplate(w, "config", guiData)
}

func saveConfigHandler(w http.ResponseWriter, r *http.Request) {

	AppConfig.Host = r.FormValue("host")
	AppConfig.Port = r.FormValue("port")
	AppConfig.Theme = r.FormValue("theme")
	AppConfig.Color = r.FormValue("color")
	AppConfig.NodePath = r.FormValue("nodepath")

	conf.Write(AppConfig.ConfPath, AppConfig)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}
