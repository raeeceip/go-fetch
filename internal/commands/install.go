package commands

import (
	"fmt"
	"log"

	"github.com/raeeceip/go-fetch/internal/core"
)

func Install(pm *core.PackageManager, args []string) {
	if len(args) != 3 {
		fmt.Println("Usage: go-fetch install <name> <version> <language>")
		return
	}

	metadata := core.PackageMetadata{
		Name:            args[0],
		Version:         args[1],
		Language:        args[2],
		MetadataVersion: "1.0.0",
	}

	err := pm.InstallPackage(metadata)
	if err != nil {
		log.Fatalf("Error installing package: %v", err)
	}
	fmt.Printf("Successfully installed %s (%s) version %s\n", metadata.Name, metadata.Language, metadata.Version)
}
