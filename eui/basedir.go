package eui

import (
	"os"
)

func getBaseDir() (string, error) {
	baseDir, _ := os.Getwd()
	return baseDir, nil
}
