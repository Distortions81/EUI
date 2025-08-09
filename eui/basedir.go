package eui

import (
	"os"
	"sync"
)

var (
	baseDir     string
	baseDirErr  error
	baseDirOnce sync.Once
)

func getBaseDir() (string, error) {
	baseDirOnce.Do(func() {
		baseDir, baseDirErr = os.Getwd()
	})
	return baseDir, baseDirErr
}
