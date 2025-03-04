package profile_test

import (
	"path/filepath"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "profile.json")
	got, err := CreateProfile(filename)
	if err != nil {
		t.Fatal(err)
	}
	want := true
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
