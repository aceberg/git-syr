package models

// Repo - git repository
type Repo struct {
	Name string `yaml:"name"`
	Data struct {
		Path    string `yaml:"path"`
		Timeout string `yaml:"timeout"`
		Pull    string `yaml:"pull"`
		Push    string `yaml:"push"`
	} `yaml:"repo"`
}
