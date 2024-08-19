package handlers

import (
	"fmt"
	"os/exec"

	"github.com/raeeceip/go-fetch/internal/core"
)

type GoHandler struct{}

func NewGoHandler() *GoHandler {
	return &GoHandler{}
}

func (h *GoHandler) Install(metadata core.PackageMetadata) error {
	fmt.Printf("Installing Go package: %s %s\n", metadata.Name, metadata.Version)
	cmd := exec.Command("go", "get", fmt.Sprintf("%s@%s", metadata.Name, metadata.Version))
	return cmd.Run()
}

func (h *GoHandler) Uninstall(metadata core.PackageMetadata) error {
	fmt.Printf("Uninstalling Go package: %s %s\n", metadata.Name, metadata.Version)
	cmd := exec.Command("go", "clean", "-i", metadata.Name)
	return cmd.Run()
}

func (h *GoHandler) List() ([]core.PackageMetadata, error) {
	cmd := exec.Command("go", "list", "-m", "all")
	_, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Parse the output and create PackageMetadata objects
	// This is a simplified version and may need to be adjusted
	var packages []core.PackageMetadata
	// ... parse output and populate packages slice

	return packages, nil
}
