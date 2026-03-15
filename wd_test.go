package wd_test

import (
	"os"
	"strings"
	"testing"

	wd "github.com/kotaoue/go-wd"
)

func TestGet(t *testing.T) {
	dir, err := wd.Get()
	if err != nil {
		t.Fatalf("Get() returned unexpected error: %v", err)
	}

	if dir == "" {
		t.Fatal("Get() returned empty string")
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Fatalf("Get() returned path that does not exist: %s", dir)
	}
}

func TestFullPath(t *testing.T) {
	const file = "testfile.txt"

	path, err := wd.FullPath(file)
	if err != nil {
		t.Fatalf("FullPath() returned unexpected error: %v", err)
	}

	if path == "" {
		t.Fatal("FullPath() returned empty string")
	}

	dir, err := wd.Get()
	if err != nil {
		t.Fatalf("Get() returned unexpected error: %v", err)
	}

	if !strings.HasPrefix(path, dir) {
		t.Fatalf("FullPath() = %q, want it to start with %q", path, dir)
	}

	if !strings.HasSuffix(path, file) {
		t.Fatalf("FullPath() = %q, want it to end with %q", path, file)
	}
}
