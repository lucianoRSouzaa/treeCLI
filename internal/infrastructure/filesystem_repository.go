package infrastructure

import (
	"os"
)

type FileSystemRepository struct{}

func NewFileSystemRepository() *FileSystemRepository {
	return &FileSystemRepository{}
}

func (fsr *FileSystemRepository) ReadDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}

	return names, nil
}

func (fsr *FileSystemRepository) IsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
