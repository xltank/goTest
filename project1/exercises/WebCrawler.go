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

type ConcMap struct {
	Map map[string]string
	mux sync.Mutex
}

func (m ConcMap) Get(k string) (string, bool) {
	m.mux.Lock()
	defer m.mux.Unlock()
	v, b := m.Map[k]
	return v, b
}

func (m ConcMap) Set(k, v string) {
	m.mux.Lock()
	defer m.mux.Unlock()
	m.Map[k] = v
}

var concMap = ConcMap{Map: make(map[string]string)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	fmt.Println("request :", url, depth)

	v, ok := concMap.Get(url)
	if ok {
		fmt.Println(url, " exists in cache, value is: ", v)
		return
	}

	concMap.Set(url, "waiting")

	body, urls, err := fetcher.Fetch(url)
	fmt.Println("body:", body, ", urls:", urls, ", err:", err)
	if err != nil {
		concMap.Set(url, fmt.Sprintln(err))
		return
	}

	concMap.Set(url, body)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
