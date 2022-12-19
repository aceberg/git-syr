package web

import (
	"bufio"
	"html/template"
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
	guiData.Icon = Icon

	file, err := os.Open(LogPath)
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

	tmpl, err := template.ParseFiles("../../web/templates/log.html", "../../web/templates/header.html", "../../web/templates/footer.html")
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "log", guiData)
	check.IfError(err)
}
