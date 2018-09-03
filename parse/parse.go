package parse

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func HTMLFiles(searchDir string) ([]string, error) {
	fileList := make([]string, 0)
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			fileList = append(fileList, path)
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return fileList, nil
}

func ExceptionFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	exceptions := strings.Split(string(content), "\n")

	return exceptions, nil
}
