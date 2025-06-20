package asset

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hellomyheart/go-indicator/helper"
)

// FileSystemRepository stores and retrieves asset snapshots using
// the local file system.
type FileSystemRepository struct {
	// base is the root directory where asset snapshots are stored.
	base string
}

// NewFileSystemRepository initializes a file system repository with
// the given base directory.
func NewFileSystemRepository(base string) *FileSystemRepository {
	return &FileSystemRepository{
		base: base,
	}
}

// Assets returns the names of all assets in the repository.
func (r *FileSystemRepository) Assets() ([]string, error) {
	files, err := os.ReadDir(r.base)
	if err != nil {
		return nil, err
	}

	var assets []string

	suffix := ".csv"

	for _, file := range files {
		name := file.Name()

		if strings.HasSuffix(name, suffix) {
			assets = append(assets, strings.TrimSuffix(name, suffix))
		}
	}

	return assets, nil
}

// Get attempts to return a channel of snapshots for the asset with the given name.
func (r *FileSystemRepository) Get(name string) (<-chan *Snapshot, error) {
	return helper.ReadFromCsvFile[Snapshot](r.getCsvFileName(name), true)
}

// GetSince attempts to return a channel of snapshots for the asset with the given name since the given date.
func (r *FileSystemRepository) GetSince(name string, date time.Time) (<-chan *Snapshot, error) {
	snapshots, err := r.Get(name)
	if err != nil {
		return nil, err
	}

	snapshots = helper.Filter(snapshots, func(s *Snapshot) bool {
		return s.Date.Equal(date) || s.Date.After(date)
	})

	return snapshots, nil
}

// LastDate returns the date of the last snapshot for the asset with the given name.
func (r *FileSystemRepository) LastDate(name string) (time.Time, error) {
	var last time.Time

	snapshots, err := r.Get(name)
	if err != nil {
		return last, err
	}

	snapshot, ok := <-helper.Last(snapshots, 1)
	if !ok {
		return last, errors.New("empty asset")
	}

	return snapshot.Date, nil
}

// Append adds the given snapshows to the asset with the given name.
func (r *FileSystemRepository) Append(name string, snapshots <-chan *Snapshot) error {
	return helper.AppendOrWriteToCsvFile(r.getCsvFileName(name), true, snapshots)
}

// getCsvFileName gets the CSV file name for the given asset name.
func (r *FileSystemRepository) getCsvFileName(name string) string {
	return filepath.Join(r.base, fmt.Sprintf("%s.csv", name))
}
