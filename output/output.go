// Package output provides all VFinder stylish outputs
package output

import (
	"fmt"
	"os"
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
	fmt.Print("\t\t|___/_/   /_/_/ /_/\\__,_/\\___/_/ (v1.2.0)\n\n")
}

func PrintFilesFound(nbFilesFound int) {
	fmt.Println(GREEN, nbFilesFound, RESET, "HTML files found")
}

// Report the analyze result
func ReportURLs(urlsError map[string]bool, urlsScanned map[string][]string) {
	fmt.Println(YELLOW, len(urlsScanned), RESET, "URLs Scanned")
	fmt.Println(RED, len(urlsError), RESET, "URLs Errored")

	if len(urlsError) != 0 {
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
		os.Exit(1)
	}
}
