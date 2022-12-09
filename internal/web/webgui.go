package web

import (
	"log"
	// "embed"
	"net/http"

	"github.com/aceberg/GitSyncTimer/internal/check"
	"github.com/aceberg/GitSyncTimer/internal/conf"
	"github.com/aceberg/GitSyncTimer/internal/models"
)

var (
	// AppConfig - config for Web Gui
	AppConfig models.Conf
	// AllRepos - all repositories
	AllRepos []models.Repo
)

// //go:embed templates/*
// var TemplHTML embed.FS

// Gui - start web server
func Gui(confPath string, allRepos []models.Repo) {

	log.Println("INFO: starting web gui with config", confPath)

	AllRepos = allRepos
	AppConfig = conf.GetConfig(confPath)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
