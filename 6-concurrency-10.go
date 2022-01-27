// https://go.dev/tour/concurrency/10

// Exercise: Web Crawler

// In this exercise you'll use Go's concurrency features to parallelize a web crawler.

// Modify the Crawl function to fetch URLs in parallel without fetching the same URL
// twice.

// Hint: you can keep a cache of the URLs that have been fetched on a map, but maps
// alone are not safe for concurrent use!

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeCache struct {
	mx   sync.Mutex
	data map[string]string
}

func (cache SafeCache) get(key string) (string, error) {
	cache.mx.Lock()
	defer cache.mx.Unlock()

	if val, ok := cache.data[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("no cache: %s", key)
}

func (cache SafeCache) set(key string, val string) {
	cache.mx.Lock()
	defer cache.mx.Unlock()

	cache.data[key] = val
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache) {

	if depth <= 0 {
		return
	}

	// Don't fetch the same URL twice.
	if _, err := cache.get(url); err == nil {
		// err is nil when the URL exists, don't do anything
		return
	}

	// TODO: Fetch URLs in parallel.

	body, urls, err := fetcher.Fetch(url)
	cache.set(url, body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher, cache)
	}

	return
}

func main() {
	cache := &SafeCache{data: make(map[string]string)}
	Crawl("https://golang.org/", 4, fetcher, cache)

	fmt.Println("-- cache --")
	for key, val := range cache.data {
		fmt.Println(key, "=", val)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
