package converter

import (
	"os"
	"path/filepath"
	"testing"
)

// createTestSVG generates a simple SVG file for testing purposes.
func createTestSVG(t *testing.T, path string) {
	content := `<svg width="100" height="100" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg"><rect width="100" height="100" fill="red"/></svg>`
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		t.Fatalf("failed to create test SVG: %v", err)
	}
}

func TestProcessSVG(t *testing.T) {
	// Setup: Create a temporary directory for test artifacts
	tmpDir := t.TempDir()
	svgPath := filepath.Join(tmpDir, "test.svg")
	outputDir := filepath.Join(tmpDir, "out")

	createTestSVG(t, svgPath)

	specs := []OutputSpec{
		{"test-16.png", 16},
		{"test-32.png", 32},
	}

	// Execute the main function
	err := ProcessSVG(svgPath, outputDir, specs, true)
	if err != nil {
		t.Fatalf("ProcessSVG failed unexpectedly: %v", err)
	}

	// Verify if expected files were created
	expectedFiles := []string{
		"test-16.png",
		"test-32.png",
		"favicon.ico",
	}

	for _, f := range expectedFiles {
		if _, err := os.Stat(filepath.Join(outputDir, f)); os.IsNotExist(err) {
			t.Errorf("expected output file missing: %s", f)
		}
	}
}
