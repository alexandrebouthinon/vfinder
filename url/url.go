package url

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"syscall"

	"golang.org/x/net/html"
)

var max_goroutines = func() uint64 {
	var rLimit syscall.Rlimit
	syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	return rLimit.Max
}

// Helper function to pull the href attribute from a Token
func getHref(t html.Token) (ok bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	return
}

// Extract all http** links from a given HTML file
func Extract(filename string) ([]string, error) {
	urls := make([]string, 0)

	c, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Failed to read %s", filename)
	}

	tkz := html.NewTokenizer(c)

	for {
		tt := tkz.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document
			return urls, nil
		case tt == html.StartTagToken:
			t := tkz.Token()

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			ok, url := getHref(t)
			if !ok {
				continue
			}

			// Make sure the url begines in http**
			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				urls = append(urls, url)
			}
		}
	}
}

// Report the analyze result
func Report(urlsError map[string]bool, urlsScanned map[string][]string) {
	fmt.Println("\033[33m", len(urlsScanned), "\033[0mURLs Scanned")
	fmt.Println("\033[31m", len(urlsError), "\033[0mURLs Errored")

	if len(urlsError) != 0 {
		fmt.Println("\n========================== \033[31mErroneous URLs\033[0m ==========================\033[31m")
		for url := range urlsError {
			fmt.Println(url, "referenced in:")
			for _, filename := range urlsScanned[url] {
				fmt.Println("\t->", filename)
			}
		}
		fmt.Println("\033[0m====================================================================")
		os.Exit(1)
	}
}

// Test all provided URLs only if they're not in excludedUrls array
func Test(urlsFoundPerFile map[string][]string, excludedUrls []string) (map[string][]string, map[string]bool) {
	var urlsScanned = make(map[string][]string)
	var urlsError = make(map[string]bool)

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	goroutines := make(chan struct{}, max_goroutines())
	for filename, urls := range urlsFoundPerFile {
		for _, url := range urls {
			if urlsScanned[url] == nil {
				urlsScanned[url] = make([]string, 1)
				urlsScanned[url][0] = filename

				goroutines <- struct{}{}
				wg.Add(1)
				go func(url, filename string) {
					_, err := http.Head(url)
					if err != nil {
						mutex.Lock()
						urlsError[url] = true
						mutex.Unlock()
					}
					<-goroutines
					wg.Done()
				}(url, filename)

			} else {
				urlsScanned[url] = append(urlsScanned[url], filename)
			}
		}
	}

	wg.Wait()

	return urlsScanned, urlsError
}

// Filter errored URLs with a given set of ignored URLs
func Filter(urlsError map[string]bool, excludedUrls []string) map[string]bool {
	for url := range urlsError {
		for _, excluded := range excludedUrls {
			if strings.HasPrefix(url, excluded) {
				delete(urlsError, url)
			}
		}
	}

	return urlsError
}
