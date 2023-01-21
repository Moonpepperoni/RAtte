package dirutil

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// GetAllFilesRecur returns the absolute path to all files in the specified dir.
// dir is searched recursively, meaning all subdirectories of all subdirectories and so on are also searched
func GetAllFilesRecur(dir string) ([]string, error) {
	var allFiles []string
	walkErr := filepath.WalkDir(dir, func(path string, pathEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !pathEntry.IsDir() {
			absolutePath, err := filepath.Abs(path)
			if err != nil {
				return err
			}
			allFiles = append(allFiles, absolutePath)
		}
		if filepath.Base(pathEntry.Name()) == "deps" {
			return filepath.SkipDir
		}
		return nil
	})
	return allFiles, walkErr
}

// FilterPaths returns all paths that match the given predicate
func FilterPaths(paths []string, predicate func(string) bool) []string {
	var filtered []string
	for _, path := range paths {
		if predicate(path) {
			filtered = append(filtered, path)
		}
	}
	return filtered
}

func relativePath(absolutePath, root string) string {
	if root == "/" {
		return absolutePath
	}
	// trimming and adding, so we know for sure cleanRoot is ending in with a slash
	cleanRoot := strings.TrimSuffix(root, "/") + "/"
	return strings.Replace(absolutePath, cleanRoot, "", 1)
}

func RelativePath(absolutePath string) string {
	wd, err := os.Getwd()
	if err != nil {
		return absolutePath
	}
	return relativePath(absolutePath, wd)

}
