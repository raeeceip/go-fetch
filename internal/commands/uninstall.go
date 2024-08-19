package commands

import (
	"fmt"
	"log"

	"github.com/raeeceip/go-fetch/internal/core"
)

func Uninstall(pm *core.PackageManager, args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: go-fetch uninstall <name> <language>")
		return
	}

	err := pm.UninstallPackage(args[0], args[1])
	if err != nil {
		log.Fatalf("Error uninstalling package: %v", err)
	}
	fmt.Printf("Successfully uninstalled %s (%s)\n", args[0], args[1])
}
