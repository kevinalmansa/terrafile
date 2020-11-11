package main

import (
	"fmt"
	"log"

	"github.com/kevinalmansa/terrafile/pkg/module"
	"github.com/spf13/viper"
)

//Config file representation
type Config struct {
	CacheDir string
	Branch   string
	Tag      string
	Modules  map[string]module.TerraformModule
}

func main() {
	var configuration Config

	viper.SetDefault("CacheDir", "modules")
	viper.SetDefault("Branch", "main")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/home/vagrant/go/src/github.com/kevinalmansa/terrafile/terrafile/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	fmt.Printf("Config: %+v\n", configuration)
}
