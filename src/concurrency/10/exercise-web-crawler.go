package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
	IsFetched(url string) bool
}

func Crawl(url string, depth int, fetcher Fetcher) {
	if fetcher.IsFetched(url) || depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	wg.Add(len(urls))
	for _, u := range urls {
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher)
		}(u)
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Second)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

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

var (
	fetchedURL = make(map[string]bool)
	mu         sync.Mutex
	wg         sync.WaitGroup
)

func (f fakeFetcher) IsFetched(url string) bool {
	mu.Lock()
	_, ok := fetchedURL[url]
	if !ok {
		fetchedURL[url] = true
	}
	mu.Unlock()
	return ok
}
