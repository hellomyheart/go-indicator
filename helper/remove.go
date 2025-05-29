package helper

import (
	"os"
	"testing"
)

// Remove删除给定名称的文件。
func Remove(t *testing.T, name string) {
	t.Helper()

	err := os.Remove(name)
	if err != nil {
		t.Errorf("Error removing file: %v", err)
	}
}

// RemoveAll删除具有给定路径的文件。
func RemoveAll(t *testing.T, path string) {
	t.Helper()

	err := os.RemoveAll(path)
	if err != nil {
		t.Errorf("Error removing files: %v", err)
	}
}
