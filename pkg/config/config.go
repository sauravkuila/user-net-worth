package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Load(configname string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(configname)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	config.AddConfigPath("../app/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}
	//for managing env values
	if configname == "server" {
		for _, v := range config.AllKeys() {
			if strings.ToLower(v) == "version" {
				continue
			}
			key := config.GetString(v)
			if key != "" {
				key = strings.ReplaceAll(key, "$", "")
				if ev, ok := os.LookupEnv(key); ok {
					config.Set(v, ev)
				} else {
					log.Fatal("env for key [", v, "] is missing")
				}
			}
		}
	}
}

func GetConfig() *viper.Viper {
	return config
}
