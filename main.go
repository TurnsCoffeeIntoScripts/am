// Copyright © 2023 Guillaume 'turnscoffeeintoscripts' Rivest <turns.coffee.into.scripts@gmail.com>
package main

import (
	"github.com/spf13/viper"
	"os"
	"turnscoffeeintoscripts/am/cmd"
	"turnscoffeeintoscripts/am/pkg/config"
)

// TODO https://zetcode.com/golang/yaml/
func main() {
	// Set proper search path for viper config
	if home := os.Getenv("HOME"); home != "" {
		viper.AddConfigPath(os.Getenv("HOME"))
	}
	viper.SetConfigName(config.FileName)
	viper.SetConfigType(config.FileType)

	// Read viper config and launch root command if successful
	//if err := viper.ReadInConfig(); err != nil {
	//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//		// TODO...
	//	} else {
	//		log.Errorf("error reading config file: %s", err)
	//	}
	//} else {
	//	cmd.Execute()
	//}

	cmd.Execute()
}
