package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListDirectoryTree(root string, ignoreList []string) (string, error) {
	var output strings.Builder
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if shouldIgnore(path, ignoreList) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		indent := strings.Repeat("  ", strings.Count(rel, string(os.PathSeparator)))
		output.WriteString(fmt.Sprintf("%s%s\n", indent, info.Name()))
		return nil
	})
	return output.String(), err
}

func shouldIgnore(path string, ignoreList []string) bool {
	for _, ignore := range ignoreList {
		if strings.Contains(path, ignore) {
			return true
		}
	}
	return false
}
