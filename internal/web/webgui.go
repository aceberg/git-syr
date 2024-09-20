package web

import (
	"log"
	"net/http"

	"github.com/aceberg/git-syr/internal/check"
	"github.com/aceberg/git-syr/internal/conf"
	"github.com/aceberg/git-syr/internal/sync"
	"github.com/aceberg/git-syr/internal/yaml"
)

// Gui - start web server
func Gui(confPath, yamlPath, logPath, nodePath string) {

	AppConfig = conf.Get(confPath)
	AppConfig.ConfPath = confPath
	AppConfig.YamlPath = yamlPath
	AppConfig.LogPath = logPath
	AppConfig.NodePath = nodePath
	AppConfig.Icon = Icon
	log.Println("INFO: starting web gui with config", AppConfig.ConfPath)

	AllRepos = yaml.Read(AppConfig.YamlPath)
	log.Println("INFO: all repos", AllRepos)

	AppConfig.Quit = make(chan bool)
	sync.AllRepos(AllRepos, AppConfig.Quit)

	address := AppConfig.Host + ":" + AppConfig.Port

	log.Println("=================================== ")
	log.Printf("Web GUI at http://%s", address)
	log.Println("=================================== ")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add_repo/", addHandler)
	http.HandleFunc("/config/", configHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/log/", logHandler)
	http.HandleFunc("/save_config/", saveConfigHandler)
	http.HandleFunc("/save_file/", saveFileHandler)
	http.HandleFunc("/update_repo/", updateHandler)
	err := http.ListenAndServe(address, nil)
	check.IfError(err)
}
