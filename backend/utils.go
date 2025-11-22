package backend

import (
	"os"
	"path/filepath"
	"strings"
)

// CountDcmFiles recursively counts all .dcm files in a folder
func CountDcmFiles(folder string) (int, error) {
	count := 0
	err := filepath.WalkDir(folder, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".dcm") {
			count++
		}
		return nil
	})
	return count, err
}
