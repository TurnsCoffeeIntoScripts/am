package config

import (
	"bufio"
	"os"
	"turnscoffeeintoscripts/am/pkg/terminal"
)

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
