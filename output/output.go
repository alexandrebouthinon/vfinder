// Package output provides all VFinder stylish outputs
package output

import (
	"fmt"
	"io/ioutil"
)

// Colors
const (
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	RESET  = "\033[0m"
)

func PrintHeader() {
	fmt.Println("\n\t\t _    _________           __")
	fmt.Println("\t\t| |  / / ____(_)___  ____/ /__  _____")
	fmt.Println("\t\t| | / / /_  / / __ \\/ __  / _ \\/ ___/")
	fmt.Println("\t\t| |/ / __/ / / / / / /_/ /  __/ /")
	fmt.Print("\t\t|___/_/   /_/_/ /_/\\__,_/\\___/_/ (v1.3.0)\n\n")
}

// Report the analyze result in STDOUT
func ReportURLsSTDOUT(nbFilesFound, nbErroredURLs, nbScannedURLs int) {
	fmt.Println(GREEN, nbFilesFound, RESET, "HTML files found")
	fmt.Println(YELLOW, nbScannedURLs, RESET, "URLs Scanned")
	fmt.Println(RED, nbErroredURLs, RESET, "URLs Errored")
}

func ReportURLsJSON(filesFound, urlsError, urlsScanned int, exportFile string) error {
	json := []byte(fmt.Sprintf(`{
		"nbFilesFound": "%d",
		"nbScannedURLs": "%d",
		"nbErroredURLs": "%d"
	}`, filesFound, urlsScanned, urlsError))

	err := ioutil.WriteFile(exportFile, json, 0666)
	if err != nil {
		return err
	}

	return nil
}

func ShowDetails(urlsScanned map[string][]string, urlsError map[string]bool) {
	fmt.Println("\n==========================", RED, "Erroneous URLs", RESET, "==========================")
	fmt.Print(RED)
	for url := range urlsError {
		fmt.Println(url, "referenced in:")
		for _, filename := range urlsScanned[url] {
			fmt.Println("\t->", filename)
		}
	}
	fmt.Print(RESET)
	fmt.Println("====================================================================")
	return
}
