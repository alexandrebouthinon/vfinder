package parse

import (
	"bufio"
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
	var exceptions []string
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			exceptions = append(exceptions, line)
		}
	}

	return exceptions, nil
}
