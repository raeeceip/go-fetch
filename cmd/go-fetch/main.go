package main

import (
	"fmt"
	"log"
	"os"

	"github.com/raeeceip/go-fetch/internal/commands"
	"github.com/raeeceip/go-fetch/internal/core"
	"github.com/raeeceip/go-fetch/internal/utils"
)

func main() {
	pm, err := core.NewPackageManager("1.0.0", "./config")
	if err != nil {
		log.Fatalf("Error creating package manager: %v", err)
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go-fetch <command> [arguments]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "install":
		commands.Install(pm, os.Args[2:])
	case "uninstall":
		commands.Uninstall(pm, os.Args[2:])
	case "list":
		commands.List(pm)
	case "tree":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go-fetch tree <directory> [ignore1,ignore2,...]")
			os.Exit(1)
		}
		ignoreList := []string{}
		if len(os.Args) > 3 {
			ignoreList = os.Args[3:]
		}
		tree, err := utils.ListDirectoryTree(os.Args[2], ignoreList)
		if err != nil {
			log.Fatalf("Error listing directory tree: %v", err)
		}
		fmt.Println(tree)
	case "mermaid":
		diagram := utils.GenerateMermaidDiagram(pm.ListPackages())
		fmt.Println(diagram)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
