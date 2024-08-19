package handlers

import (
	"fmt"
	"os/exec"

	"github.com/raeeceip/go-fetch/internal/core"
)

type RubyHandler struct{}

func NewRubyHandler() *RubyHandler {
	return &RubyHandler{}
}

func (h *RubyHandler) Install(metadata core.PackageMetadata) error {
	fmt.Printf("Installing Ruby package: %s %s\n", metadata.Name, metadata.Version)
	cmd := exec.Command("gem", "install", metadata.Name, "-v", metadata.Version)
	return cmd.Run()
}

func (h *RubyHandler) Uninstall(metadata core.PackageMetadata) error {
	fmt.Printf("Uninstalling Ruby package: %s %s\n", metadata.Name, metadata.Version)
	cmd := exec.Command("gem", "uninstall", metadata.Name, "-v", metadata.Version)
	return cmd.Run()
}

func (h *RubyHandler) List() ([]core.PackageMetadata, error) {
	cmd := exec.Command("gem", "list", "--local")
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
