package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// GetCurrFileDir returns the directory of the caller
func GetCurrFileDir() string {
	// get this file path
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(fmt.Errorf("Unable to access the current file path"))
	}
	return filepath.Dir(file)
}
