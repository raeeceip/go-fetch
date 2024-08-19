package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Masterminds/semver/v3"
	"github.com/raeeceip/go-fetch/internal/handlers"
)

type PackageManager struct {
	version   *semver.Version
	handlers  map[string]LanguageHandler
	packages  []PackageMetadata
	configDir string
}

type LanguageHandler interface {
	Install(metadata PackageMetadata) error
	Uninstall(metadata PackageMetadata) error
	List() ([]PackageMetadata, error)
}

func NewPackageManager(version string, configDir string) (*PackageManager, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return nil, err
	}

	pm := &PackageManager{
		version:   v,
		handlers:  make(map[string]LanguageHandler),
		packages:  []PackageMetadata{},
		configDir: configDir,
	}

	// Register language handlers
	pm.handlers["go"] = handlers.NewGoHandler()
	pm.handlers["ruby"] = handlers.NewRubyHandler()

	err = pm.loadPackages()
	if err != nil {
		return nil, err
	}

	return pm, nil
}

func (pm *PackageManager) loadPackages() error {
	data, err := ioutil.ReadFile(filepath.Join(pm.configDir, "packages.json"))
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	return json.Unmarshal(data, &pm.packages)
}

func (pm *PackageManager) savePackages() error {
	data, err := json.MarshalIndent(pm.packages, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(pm.configDir, "packages.json"), data, 0644)
}

func (pm *PackageManager) InstallPackage(metadata PackageMetadata) error {
	handler, ok := pm.handlers[metadata.Language]
	if !ok {
		return fmt.Errorf("unsupported language: %s", metadata.Language)
	}

	err := handler.Install(metadata)
	if err != nil {
		return err
	}

	pm.packages = append(pm.packages, metadata)
	return pm.savePackages()
}

func (pm *PackageManager) UninstallPackage(name string, language string) error {
	handler, ok := pm.handlers[language]
	if !ok {
		return fmt.Errorf("unsupported language: %s", language)
	}

	var metadata PackageMetadata
	var index int
	for i, pkg := range pm.packages {
		if pkg.Name == name && pkg.Language == language {
			metadata = pkg
			index = i
			break
		}
	}

	if metadata.Name == "" {
		return fmt.Errorf("package not found: %s (%s)", name, language)
	}

	err := handler.Uninstall(metadata)
	if err != nil {
		return err
	}

	pm.packages = append(pm.packages[:index], pm.packages[index+1:]...)
	return pm.savePackages()
}

func (pm *PackageManager) ListPackages() []PackageMetadata {
	return pm.packages
}
