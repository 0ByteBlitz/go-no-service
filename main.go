package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

// ProxyHandler handles requests and forwards them to the target URL
type ProxyHandler struct {
	TargetURL *url.URL
}

func (p *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Modify the request's host to target the desired URL
	r.Host = p.TargetURL.Host
	// Reverse proxy the request to the target URL
	proxy := httputil.NewSingleHostReverseProxy(p.TargetURL)
	proxy.ServeHTTP(w, r)
}

// function to make a specified number of requests to the proxy server
func makeConcurrentRequest(proxyURL string, targetURL string, numRequests int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	proxyURLParsed, err := url.Parse(proxyURL)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v", err)
		return
	}

	for i := 0; i < numRequests; i++ {
		client := &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyURLParsed),
			},
		}
		start := time.Now()
		resp, err := client.Get(targetURL)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
			return
		}
		secs := time.Since(start).Seconds()
		ch <- fmt.Sprintf("Success: %s %d %v", targetURL, resp.StatusCode, secs)
	}
}

// main function
func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: ./main.go <proxyURL> <targetURL> <numRequests>")
		return
	}

	proxyURL := os.Args[1]
	targetURL := os.Args[2]
	numRequests, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid numRequests argument:", err)
		return
	}

	ch := make(chan string)
	var wg sync.WaitGroup

	fmt.Printf("Making %d concurrent requests to the target URL: %s\n", numRequests, targetURL)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go makeConcurrentRequest(proxyURL, targetURL, numRequests, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
}
