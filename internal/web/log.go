package web

import (
	"bufio"
	"net/http"
	"os"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var text string

	guiData.Config = AppConfig
	guiData.Icon = Icon

	file, err := os.Open(LogPath)
	check.IfError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = scanner.Text()
		guiData.Themes = append(guiData.Themes, text)
	}

	file.Close()

	execTemplate(w, "log", guiData)
}
