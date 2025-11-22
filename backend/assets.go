package backend

import (
	"embed"
	"io/fs"
	"os"
	"path/filepath"
)

// Embed the entire dcm4che-5.34.1 folder
//
//go:embed dcm4che-5.34.1/**/*
var dcm4cheFiles embed.FS

// ExtractAssets extracts embedded dcm4che to a temp folder
func ExtractAssets() (string, error) {
	tempDir, err := os.MkdirTemp("", "dcm4che")
	if err != nil {
		return "", err
	}

	// Copy all files
	err = fs.WalkDir(dcm4cheFiles, "dcm4che-5.34.1", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		data, err := dcm4cheFiles.ReadFile(path)
		if err != nil {
			return err
		}
		dest := filepath.Join(tempDir, filepath.Base(path))
		return os.WriteFile(dest, data, 0755)
	})
	if err != nil {
		return "", err
	}

	return tempDir, nil
}
