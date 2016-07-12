package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCache struct {	// map caching the url found, added by me
	cache map[string]int
	mux sync.Mutex
}

func Yell(){
	fmt.Println("hi")
}
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	urlPot.mux.Lock()
	urlPot.cache[url] = 1
	urlPot.mux.Unlock()
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	//fmt.Println(urls)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		urlPot.mux.Lock() // check it's is old url or not
		_,ok := urlPot.cache[u]
		if !ok {
			urlPot.cache[u] = 1
			//Crawl(u, depth-1, fetcher)
		}
		urlPot.mux.Unlock()
		if !ok && u !=url {
			//fmt.Println(url,"->",u)
			go Crawl(u, depth-1, fetcher)
		}
	}
	return
}

func main() {
	urlPot = UrlCache{cache: make(map[string]int)}
	//Crawl("http://golang.org/", 4, fetcher)
	Crawl("http://golang.org/", 4, fetcher)
	time.Sleep(1000 * time.Millisecond)
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

var urlPot UrlCache
