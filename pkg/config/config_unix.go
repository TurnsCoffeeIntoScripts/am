package config

import (
	"bufio"
	"bytes"
	"os"
	"turnscoffeeintoscripts/am/pkg/model"
	"turnscoffeeintoscripts/am/pkg/terminal"

	"github.com/spf13/viper"
)

func Setup() {
	home := os.Getenv("HOME")
	if home != "" {
		viper.AddConfigPath(home)
	}
	viper.SetConfigName(FileName)
	viper.SetConfigType(FileType)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			initConfigFile(home)
		}
	}
}

func Exists() bool {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return false
		}
	}

	return true
}

func initConfigFile(home string) {
	if _, err := os.Create(home + "/" + FileName + "." + FileType); err != nil {
		terminal.ErrorMessage("Failed to create config file (%s)", err.Error())
	} else {
		s := Storage{}

		temp := model.Alias{
			Name:     "ll",
			Command:  "ls -la",
			Category: "none",
		}
		s.Aliases = append(s.Aliases, temp)
		viper.ReadConfig(bytes.NewBuffer(s.Marshal()))
		if err := viper.WriteConfig(); err != nil {
			terminal.ErrorMessage("Failed to init config file (%s)", err.Error())
		}
	}
}

// ===========================================================
// ===========================================================
// OLD STUFF BELOW
// ===========================================================
// ===========================================================
func AMConfigExists() (bool, string) {
	return fileExists(FileName + "." + FileType)
}

func AMFileExists() (bool, string) {
	return fileExists(".am")
}

func CreateAMConfig() string {
	home := getHome()
	f, err := os.Create(home + "/" + FileName + "." + FileType)
	if err != nil {
		terminal.ErrorMessage("Failed to create %s.%s", FileName, FileType)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			terminal.ErrorMessage("Failed to close %s.%s file handler", FileName, FileType)
		}
	}(f)

	// Write default am.yaml content
	writeDefaultAMConfig(f)

	return f.Name()
}

func fileExists(file string) (bool, string) {
	var info os.FileInfo
	var err error
	home := getHome()
	info, err = os.Stat(home + "/" + file)

	if os.IsNotExist(err) {
		return false, home
	}

	return !info.IsDir(), home
}

func getHome() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	} else {
		return "~"
	}
}

func writeDefaultAMConfig(f *os.File) {
	w := bufio.NewWriter(f)
	_, err := w.WriteString("aliases:")
	if err != nil {
		terminal.ErrorMessage("Failed to write to am.yaml")
		panic("Failed to write to am.yaml")
	}
	_ = w.Flush()
}
