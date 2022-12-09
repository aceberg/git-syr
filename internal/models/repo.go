package models

// Repo - git repository
type Repo struct {
	ID      int    `yaml:"id"`
	Name    string `yaml:"name"`
	Path    string `yaml:"path"`
	Timeout string `yaml:"timeout"`
	Pull    string `yaml:"pull"`
	Push    string `yaml:"push"`
}
