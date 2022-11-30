package web

import (
	// "fmt"
	// "embed"
	. "github.com/aceberg/GitSyncTimer/internal/models"
	"log"
	// "net/http"
)

// var AppConfig Conf
// var TableList []Table

// //go:embed templates/*
// var TemplHTML embed.FS

func Gui(confPath string, allRepos []Repo) {

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
