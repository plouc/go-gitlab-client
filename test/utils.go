package test

import (
	"log"
	"path/filepath"
	"runtime"
	"testing"
)

func baseDir(t *testing.T) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		if t != nil {
			t.Fatal("an error occurred while recovering caller information")
			return ""
		} else {
			log.Fatal("an error occurred while recovering caller information")
		}
	}

	return filepath.Dir(filename)
}
