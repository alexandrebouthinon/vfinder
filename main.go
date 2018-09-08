package main

import (
	"flag"
	"os"

	"github.com/alexandrebouthinon/vfinder/output"
	"github.com/alexandrebouthinon/vfinder/parse"
	"github.com/alexandrebouthinon/vfinder/url"
)

var (
	directoryRoot *string
	targetFile    *string
	exceptionFile *string
)

func init() {
	output.PrintHeader()
	directoryRoot = flag.String("d", "", "A directory location as a string, this directory or sub-directories should contain HTML files to analize.")
	targetFile = flag.String("f", "", "A file path as a string, This file should contain HTML code.")
	exceptionFile = flag.String("x", "", "An exception filename as a string, this file sould contains prefix that need to be ignored in parsing.")
}

func main() {
	var urlsFoundPerFile = make(map[string][]string)
	var excludedUrls = make([]string, 0)
	var err error

	flag.Parse()

	if (*directoryRoot == "" && *targetFile == "") || (*directoryRoot != "" && *targetFile != "") {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *exceptionFile != "" {
		excludedUrls, err = parse.ExceptionFile(*exceptionFile)
		if err != nil {
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	if *directoryRoot != "" {
		filesFound, err := parse.HTMLFiles(*directoryRoot)
		if err != nil {
			panic(err)
		}
		output.PrintFilesFound(len(filesFound))
		for _, f := range filesFound {
			urls, err := url.Extract(f)
			if err != nil {
				panic(err)
			}
			urlsFoundPerFile[f] = urls
		}
	} else if *targetFile != "" {
		urls, err := url.Extract(*targetFile)
		if err != nil {
			panic(err)
		}
		urlsFoundPerFile[*targetFile] = urls
	}

	// Urls testing
	urlsScanned, urlsError := url.Test(urlsFoundPerFile, excludedUrls)

	if len(excludedUrls) != 0 {
		urlsError = url.Filter(urlsError, excludedUrls)
	}

	// Scan report
	output.ReportURLs(urlsError, urlsScanned)
}
