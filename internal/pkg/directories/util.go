package directories

import (
	"os"
	"path/filepath"
)

func FindDirectoryName(directoryName string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var searchDir func(string) string
	searchDir = func(currentDir string) string {
		if filepath.Base(currentDir) == directoryName {
			return currentDir
		}

		entries, err := os.ReadDir(currentDir)
		if err != nil {
			panic(err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				found := searchDir(filepath.Join(currentDir, entry.Name()))
				if found != "" {
					return found
				}
			}
		}

		return ""
	}

	for {
		found := searchDir(dir)
		if found != "" {
			return found
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return ""
}
