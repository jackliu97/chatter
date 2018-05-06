package config

import (
	"github.com/spf13/viper"
	"fmt"
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
