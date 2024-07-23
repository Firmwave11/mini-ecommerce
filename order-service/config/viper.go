package config

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Load() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// Set default
	viper.SetDefault("application.port", 8080)
	viper.SetDefault("db.maxidle", 5)
	viper.SetDefault("db.maxconn", 10)
	viper.SetDefault("cors.whitelist", []string{"*"})

	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error(err)
	}
}
