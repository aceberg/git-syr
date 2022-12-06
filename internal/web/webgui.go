package web

import (
	"log"
	// "fmt"
	// "embed"
	// "net/http"

	"github.com/aceberg/GitSyncTimer/internal/models"
)

// var AppConfig Conf
// var TableList []Table

// //go:embed templates/*
// var TemplHTML embed.FS

// Gui - start web server
func Gui(confPath string, allRepos []models.Repo) {

	log.Println("INFO: starting web gui with config", confPath)

	// AppConfig = appConfig
	// TableList = db.SelectTableList(AppConfig.DbPath)

	// address := AppConfig.GuiIP + ":" + AppConfig.GuiPort

	// log.Println("=================================== ")
	// log.Printf("Web GUI at http://%s", address)
	// log.Println("=================================== ")

	// http.HandleFunc("/", dashboard)
	// http.ListenAndServe(address, nil)
}
