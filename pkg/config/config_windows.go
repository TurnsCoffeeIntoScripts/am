package config

import (
	"os"
)

func AMFileExists() (bool, string) {
	var info os.FileInfo
	var err error
	var location string
	if home := os.Getenv("HOME"); home != "" {
		location = home
		info, err = os.Stat(home + "\\" + FileName + "." + FileType)
	} else {
		// TODO find proper equivalent for '~/' for windows...
	}

	if os.IsNotExist(err) {
		return false, location
	}

	return !info.IsDir(), location
}
