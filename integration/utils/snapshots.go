package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

type SnapshotFile struct {
	T    *testing.T
	Name string
	Dir  string
}

func NewSnapshotFile(t *testing.T, name, dir string) *SnapshotFile {
	return &SnapshotFile{
		T:    t,
		Name: name + ".snapshot",
		Dir:  dir,
	}
}

func (s *SnapshotFile) Path() string {
	return filepath.Join(s.Dir, s.Name)
}

func (s *SnapshotFile) Load() string {
	s.T.Helper()

	content, err := ioutil.ReadFile(s.Path())
	if err != nil {
		s.T.Fatalf("could not read file %s: %v", s.Name, err)
	}

	return string(content)
}

func (s *SnapshotFile) Write(content string) {
	s.T.Helper()
	err := ioutil.WriteFile(s.Path(), []byte(content), 0644)
	if err != nil {
		s.T.Fatalf("could not write %s: %v", s.Name, err)
	}
}

func (s *SnapshotFile) File() *os.File {
	s.T.Helper()
	file, err := os.Open(s.Path())
	if err != nil {
		s.T.Fatalf("could not open %s: %v", s.Name, err)
	}

	return file
}
