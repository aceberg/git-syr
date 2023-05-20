package web

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

func errorLogHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var text string

	guiData.Config = AppConfig

	file, err := os.Open(AppConfig.LogPath)
	check.IfError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = scanner.Text()
		if strings.Contains(text, "ERROR") {
			guiData.Themes = append(guiData.Themes, text)
		}
	}

	file.Close()

	execTemplate(w, "log", guiData)
}
