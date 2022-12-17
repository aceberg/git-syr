package web

import (
	"log"
	// "embed"
	"net/http"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/conf"
	"github.com/aceberg/git-syr/internal/models"
)

var (
	// AppConfig - config for Web Gui
	AppConfig models.Conf
	// AllRepos - all repositories
	AllRepos []models.Repo
	// ConfigPath - path to Gui config file
	ConfigPath string
	// YamlPath - path to repos file
	YamlPath string
	// LogPath - path to log file
	LogPath string
)

// //go:embed templates/*
// var TemplHTML embed.FS

// Gui - start web server
func Gui(confPath, yamlPath, logPath string, allRepos []models.Repo) {

	AllRepos = allRepos
	ConfigPath = confPath
	YamlPath = yamlPath
	LogPath = logPath

	log.Println("INFO: starting web gui with config", ConfigPath)
	AppConfig = conf.Get(ConfigPath)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add_repo/", addHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/log/", logHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/update_repo/", updateHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
