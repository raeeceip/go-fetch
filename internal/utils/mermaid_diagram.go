package utils

import (
	"fmt"
	"strings"

	"github.com/raeeceip/go-fetch/internal/core"
)

func GenerateMermaidDiagram(packages []core.PackageMetadata) string {
	var diagram strings.Builder
	diagram.WriteString("graph TD\n")
	for _, pkg := range packages {
		diagram.WriteString(fmt.Sprintf("  %s[%s %s]\n", pkg.Name, pkg.Name, pkg.Version))
	}
	// Add some example relationships
	for i := 0; i < len(packages)-1; i++ {
		diagram.WriteString(fmt.Sprintf("  %s --> %s\n", packages[i].Name, packages[i+1].Name))
	}
	return diagram.String()
}
