package conf

import (
	. "github.com/aceberg/GitSyncTimer/internal/common"
	. "github.com/aceberg/GitSyncTimer/internal/models"
	"github.com/spf13/viper"
)

func GetConfig(configPath string) (config Conf) {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	CheckIfError(err)

	return config
}

func WriteConfig(configPath string, config Conf) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.Set("THEME", theme)
	viper.WriteConfig()
}
