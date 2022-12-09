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
	// ConfigPath - path to Gui config file
	ConfigPath string
)

// //go:embed templates/*
// var TemplHTML embed.FS

// Gui - start web server
func Gui(confPath string, allRepos []models.Repo) {

	AllRepos = allRepos
	ConfigPath = confPath

	log.Println("INFO: starting web gui with config", ConfigPath)
	AppConfig = conf.Get(ConfigPath)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/add_repo/", addHandler)
	http.HandleFunc("/config/", configHandler)
	// http.HandleFunc("/log/", logHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	// http.HandleFunc("/update_repo/", updateHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
