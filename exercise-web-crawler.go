// concurrency/10
// Even thought I 'stole' only half of a code at
// https://rmoff.net/2020/07/03/learning-golang-some-rough-notes-s01e10-concurrency-web-crawler/
// and before that read bunch of other hints and solutions,
// in the end solution is no different from that one so credit goes there.

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

type Fetched struct {
	mu sync.Mutex
	v  map[string]bool
}

func (f Fetched) Get (url string) bool {
	// Check was url crawled
	f.mu.Lock()
	defer f.mu.Unlock()
	_, ok := f.v[url]
	return ok
}

func (f Fetched) Set (url string) {
	f.mu.Lock()
	f.v[url] = true
	f.mu.Unlock()
}

var fetched = Fetched{v: make(map[string]bool)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()
	
	if fetched.Get(url) == true {
		return
	}
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	
	fetched.Set(url)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg)
	}
	return
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, wg)
	wg.Wait()
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
