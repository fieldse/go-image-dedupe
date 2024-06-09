// The CLI for the program
package cli

import (
	"fmt"

	_ "github.com/fieldse/go-image-dedupe/internal/dedupe"
	_ "github.com/fieldse/go-image-dedupe/internal/scanner"
	"github.com/fieldse/go-image-dedupe/internal/shared"
)

func RunCLI() {}

// getImageDirs prompts the user for target image directories to scan
func getImageDirs() ([]string, error) {
	return []string{}, fmt.Errorf("not implemented")
}

// getDestDir prompts the user for destination directory to organize to
func getDestDir() ([]string, error) {
	return []string{}, fmt.Errorf("not implemented")
}

// getScanMode prompts the user for scan mode
func getScanMode() (shared.ScanType, error) {
	return shared.ScanTypeNone, fmt.Errorf("not implemented")
}
