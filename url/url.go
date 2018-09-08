package url

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

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
			if isAnchor := t.Data == "a"; !isAnchor {
				continue
			}
			// Extract the href value, if there is one
			ok, url := getHref(t)
			if !ok {
				continue
			}
			// Make sure the url begines in http**
			if strings.Index(url, "http") == 0 {
				urls = append(urls, url)
			}
		}
	}
}

// Test all provided URLs only if they're not in excludedUrls array
func Test(urlsFoundPerFile map[string][]string, excludedUrls []string) (map[string][]string, map[string]bool) {
	var urlsScanned = make(map[string][]string)
	var urlsError = make(map[string]bool)

	var wg sync.WaitGroup
	var mutex = &sync.Mutex{}
	goroutines := make(chan struct{}, 100)
	for filename, urls := range urlsFoundPerFile {
		for _, url := range urls {
			if urlsScanned[url] == nil {
				urlsScanned[url] = make([]string, 1)
				urlsScanned[url][0] = filename

				goroutines <- struct{}{}
				wg.Add(1)
				go func(url, filename string) {
					if _, err := http.Head(url); err != nil {
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
	close(goroutines)

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
