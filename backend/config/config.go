package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Env string

const (
	DevelopmentEnv Env = "development"
	StagingEnv     Env = "staging"
	ProductionEnv  Env = "production"
)

type Config struct {
	BaseUrl            string
	DatabaseUrl        string
	DatabaseName       string
	DatabaseInitialCap int
	DatabaseMaxOpen    int
}

func ReadConfig(env Env) Config {
	viper.SetConfigName("app")
	viper.AddConfigPath("./config")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	baseUrl := fmt.Sprintf("%s.%s", env, "BaseUrl")
	databaseUrl := fmt.Sprintf("%s.%s", env, "DatabaseUrl")
	databaseName := fmt.Sprintf("%s.%s", env, "DatabaseName")
	initialCap := fmt.Sprintf("%s.%s", env, "DatabaseInitialCap")
	maxOpen := fmt.Sprintf("%s.%s", env, "DatabaseMaxOpen")

	return Config{
		viper.GetString(baseUrl),
		viper.GetString(databaseUrl),
		viper.GetString(databaseName),
		viper.GetInt(initialCap),
		viper.GetInt(maxOpen),
	}
}
