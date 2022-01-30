package wd

import (
	"os"
	"path/filepath"
)

func Get() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	if filepath.Base(exe) == "main" {
		return os.Getwd()
	}
	return filepath.Dir(exe), nil
}
