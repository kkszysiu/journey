package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

func GetFilenameWithoutExtension(path string) string {
	partials_path := string(filepath.Separator) + "partials" + string(filepath.Separator)
	filename := ""

	if strings.Contains(path, partials_path) {
		filename = path[strings.Index(path, partials_path)+len(partials_path) : len(path)-len(filepath.Ext(path))]
	} else {
		filename = filepath.Base(path)[0 : len(filepath.Base(path))-len(filepath.Ext(path))]
	}

	return filename
}

func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		// We could check with os.IsNotExist(err) here, but since os.Stat threw an error, we likely can't use the file anyway.
		return false
	}
	return true
}
