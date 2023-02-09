package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"os"
	"fmt"
	"time"
)

func main() {
	
	warning()

	urls := []string{
		"https://www.353solutions.com/c/znga/",
		"https://www.deepl.com/translator",
		"https://www.zhihu.com/",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		// wg.Add(1)
		url := url
		go func() {
			defer wg.Done()
			siteTime(url)
		}()
	}
	
	wg.Wait()
}

func siteTime(url string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("ERROR: %s -> %s", url, err)
	}

	duration := time.Since(start)
	log.Printf("INFO: %s -> %v", url, duration)
}





func warning() {
	fmt.Fprintln(os.Stderr, "┌──────────────────────────────────────────────────────┐")
	fmt.Fprintln(os.Stderr, "│                                                      │")
	fmt.Fprintln(os.Stderr, "│   please visit:                                      │")
	fmt.Fprintln(os.Stderr, "│         https://www.sjxiang.com/                     │")
	fmt.Fprintln(os.Stderr, "│                                                      │")
	fmt.Fprintln(os.Stderr, "└──—————————————————————————————————————————————————───┘")
}