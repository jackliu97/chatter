package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	fmt.Println("init Config...")

	viper.SetConfigType("yaml")
	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
