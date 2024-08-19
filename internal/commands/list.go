package commands

import (
	"fmt"

	"github.com/raeeceip/go-fetch/internal/core"
)

func List(pm *core.PackageManager) {
	packages := pm.ListPackages()
	if len(packages) == 0 {
		fmt.Println("No packages installed.")
		return
	}

	fmt.Println("Installed packages:")
	for _, pkg := range packages {
		fmt.Printf("- %s (%s) version %s\n", pkg.Name, pkg.Language, pkg.Version)
	}
}
