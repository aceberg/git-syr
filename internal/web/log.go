package web

import (
	"bufio"
	"html/template"
	"net/http"
	"os"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/models"
)

func logHandler(w http.ResponseWriter, r *http.Request) {
	var guiData models.GuiData
	var text []string

	guiData.Config = AppConfig
	guiData.Icon = Icon

	file, err := os.Open(LogPath)
	check.IfError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	guiData.Themes = text

	tmpl, err := template.ParseFiles("../../web/templates/log.html", "../../web/templates/header.html", "../../web/templates/footer.html")
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "header", guiData)
	check.IfError(err)
	err = tmpl.ExecuteTemplate(w, "log", guiData)
	check.IfError(err)
}
