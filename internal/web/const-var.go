package web

import (
	"embed"

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
	// Quit - channel
	Quit chan bool
)

// TemplHTML - html templates
//
//go:embed templates/*
var TemplHTML embed.FS

// TemplPath - path to html templates
const TemplPath = "templates/"
