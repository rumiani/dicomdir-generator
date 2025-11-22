package backend

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

// extractDcm4Che extracts the embedded dcm4che folder to a temp directory
func extractDcm4Che() (string, error) {
	tempDir, err := os.MkdirTemp("", "dcm4che")
	if err != nil {
		return "", err
	}

	err = fs.WalkDir(dcm4cheFiles, "dcm4che-5.34.1", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		data, err := dcm4cheFiles.ReadFile(path)
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel("dcm4che-5.34.1", path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(tempDir, relPath)
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return err
		}

		return os.WriteFile(destPath, data, 0755)
	})
	if err != nil {
		return "", err
	}

	return tempDir, nil
}

// GenerateDicomDir generates a DICOMDIR inside the selected folder
func GenerateDicomDir(folder string) error {
	// Extract embedded dcm4che to temp folder
	tempDcm4Che, err := extractDcm4Che()
	if err != nil {
		return fmt.Errorf("failed to extract dcm4che: %w", err)
	}

	// Path to dcmdir executable
	dcmdirPath := filepath.Join(tempDcm4Che, "bin", "dcmdir")

	// Ensure dcmdir is executable (especially on Linux/macOS)
	if err := os.Chmod(dcmdirPath, 0755); err != nil {
		return fmt.Errorf("failed to set executable permission on dcmdir: %w", err)
	}

	// Build command: dcmdir -c <output DICOMDIR path> <input folder>
	cmd := exec.Command(
		dcmdirPath,
		"-c",
		filepath.Join(folder, "DICOMDIR"), // output DICOMDIR inside folder
		folder,                            // input folder
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run dcmdir: %w", err)
	}

	return nil
}
