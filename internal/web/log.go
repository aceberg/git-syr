package web

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var text string

	guiData.Config = AppConfig

	mode := r.URL.Query().Get("mode")

	file, err := os.Open(AppConfig.LogPath)
	check.IfError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = scanner.Text()
		if mode == "error" && strings.Contains(text, "ERROR") {
			guiData.Themes = append(guiData.Themes, text)
		} else if mode != "error" {
			guiData.Themes = append(guiData.Themes, text)
		}
	}

	file.Close()

	if mode != "all" {
		logLength := len(guiData.Themes)
		if logLength > 20 {
			guiData.Themes = guiData.Themes[logLength-20 : logLength]
		}
	}

	execTemplate(w, "log", guiData)
}
