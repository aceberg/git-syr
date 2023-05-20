package web

import (
	"net/http"

	"github.com/aceberg/git-syr/internal/models"
)

func helpHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData

	guiData.Config = AppConfig

	guiData.Themes = []string{"● To delete repository save it with empty path",
		"● Timeout examples: '30s' - 30 seconds, '5m' - 5 minutes, '3h' - 3 hours"}

	execTemplate(w, "log", guiData)
}
