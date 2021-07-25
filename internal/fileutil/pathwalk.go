package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
)

//PathWalk will walk up the path until it finds what you're looking for
func PathWalk(folderPath string) (string, error) {
	targetPath := folderPath

	var abs string

	for {

		newAbs, _ := filepath.Abs(targetPath)
		if newAbs == abs {
			return "", fmt.Errorf(fmt.Sprintf("Could not find path %v", folderPath))
		}

		abs = newAbs

		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			targetPath = filepath.Join("..", targetPath)
			continue
		}

		return targetPath, nil
	}
}

//FileWalk will walk up the path until it finds the file you're looking for
func FileWalk(folderPath string, fileName string) (string, error) {
	targetPath := folderPath

	for {

		searchFile := filepath.Join(targetPath, fileName)

		if _, err := os.Stat(searchFile); os.IsNotExist(err) {
			targetPath = filepath.Dir(targetPath)
			if targetPath == "/" {
				return "", fmt.Errorf(fmt.Sprintf("Could not find path %v", folderPath))
			}
			continue
		}

		return targetPath, nil
	}
}
