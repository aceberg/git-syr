package models

// type Repo struct {
// 	Path    string `yaml:"Path"`
// 	Timeout string `yaml:"Timeout"`
// 	Pull    string `yaml:"Pull"`
// 	Push    string `yaml:"Push"`
// }

type OneRepo struct {
	Name string `yaml:"name"`
	Repo struct {
		Path    string `yaml:"path"`
		Timeout string `yaml:"timeout"`
		Pull    string `yaml:"pull"`
		Push    string `yaml:"push"`
	} `yaml:"repo"`
}
