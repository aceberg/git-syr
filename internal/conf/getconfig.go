package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/GitSyncTimer/internal/check"
	"github.com/aceberg/GitSyncTimer/internal/models"
)

// GetConfig - read config from file or env
func GetConfig(path string) models.Conf {
	var config models.Conf

	viper.SetDefault("HOST", "0.0.0.0")
	viper.SetDefault("PORT", "8844")
	viper.SetDefault("THEME", "minty")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	check.IfError(err)

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.Host, _ = viper.Get("HOST").(string)
	config.Port, _ = viper.Get("PORT").(string)
	config.Theme, _ = viper.Get("THEME").(string)

	return config
}

// func WriteConfig(theme string) {
// 	viper.SetConfigFile(configPath)
// 	viper.SetConfigType("env")
// 	viper.Set("THEME", theme)
// 	viper.WriteConfig()
// }
