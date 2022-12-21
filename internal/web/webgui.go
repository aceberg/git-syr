package web

import (
	"embed"
	"log"
	"net/http"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/conf"
	"github.com/aceberg/git-syr/internal/models"
	"github.com/aceberg/git-syr/internal/sync"
	"github.com/aceberg/git-syr/internal/yaml"
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
	// Quit - channel
	Quit chan bool
)

// TemplHTML - html templates
//
//go:embed templates/*
var TemplHTML embed.FS

// Gui - start web server
func Gui(confPath, yamlPath, logPath string) {

	// AllRepos = allRepos
	ConfigPath = confPath
	YamlPath = yamlPath
	LogPath = logPath

	AllRepos = yaml.Read(YamlPath)
	log.Println("INFO: all repos", AllRepos)

	Quit = make(chan bool)
	sync.AllRepos(AllRepos, Quit)

	log.Println("INFO: starting web gui with config", ConfigPath)
	AppConfig = conf.Get(ConfigPath)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add_repo/", addHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/error/", errorLogHandler)
	http.HandleFunc("/help/", helpHandler)
	http.HandleFunc("/log/", logHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/update_repo/", updateHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
