package models

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Icon   string
	Repos  []Repo
	Themes []string
}
