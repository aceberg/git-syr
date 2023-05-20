package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Icon     string
	ConfPath string
	YamlPath string
	LogPath  string
	Quit     chan bool
}

// Repo - git repository
type Repo struct {
	ID      int    `yaml:"id"`
	Name    string `yaml:"name"`
	Path    string `yaml:"path"`
	Timeout string `yaml:"timeout"`
	Pull    string `yaml:"pull"`
	Push    string `yaml:"push"`
	AddPush string `yaml:"addpush"`
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Icon   string
	Repos  []Repo
	Themes []string
}
