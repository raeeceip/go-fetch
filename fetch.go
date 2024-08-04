package main

import (
	"fmt"
	"errors"

	"github.com/Masterminds/semver/v3"
)

type PackageMetadata struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	MetadataVersion string `json:"metadata_version"`
}

type PackageManagerPlugin interface {
	CanHandle(metadata PackageMetadata) bool
	Install(metadata PackageMetadata) error
}

type PackageManager struct {
	version *semver.Version
	plugins []PackageManagerPlugin
}

func NewPackageManager(version string) (*PackageManager, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return nil, err
	}
	return &PackageManager{
		version: v,
		plugins: make([]PackageManagerPlugin, 0),
	}, nil
}

func (pm *PackageManager) RegisterPlugin(plugin PackageManagerPlugin) {
	pm.plugins = append(pm.plugins, plugin)
}

func (pm *PackageManager) InstallPackage(metadata PackageMetadata) error {
	metadataVersion, err := semver.NewVersion(metadata.MetadataVersion)
	if err != nil {
		return err
	}

	if metadataVersion.GreaterThan(pm.version) {
		return fmt.Errorf("package metadata version %s is not supported", metadataVersion)
	}

	for _, plugin := range pm.plugins {
		if plugin.CanHandle(metadata) {
			return plugin.Install(metadata)
		}
	}

	return errors.New("no plugin can handle this package")
}

type BasicPackageHandler struct{}

func (h *BasicPackageHandler) CanHandle(metadata PackageMetadata) bool {
	v, err := semver.NewVersion(metadata.MetadataVersion)
	if err != nil {
		return false
	}
	return v.LessThanOrEqual(semver.MustParse("1.0.0"))
}

func (h *BasicPackageHandler) Install(metadata PackageMetadata) error {
	fmt.Printf("Installing %s version %s\n", metadata.Name, metadata.Version)
	return nil
}

func main() {
	pm, err := NewPackageManager("1.0.0")
	if err != nil {
		fmt.Println("Error creating package manager:", err)
		return
	}

	pm.RegisterPlugin(&BasicPackageHandler{})

	metadata := PackageMetadata{
		Name:            "example-package",
		Version:         "0.1.0",
		MetadataVersion: "1.0.0",
	}

	err = pm.InstallPackage(metadata)
	if err != nil {
		fmt.Println("Error installing package:", err)
	}
}
