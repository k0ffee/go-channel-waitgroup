package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"sync"
)

type respHandler struct {
	mu sync.Mutex
	n  uint64
}

var (
	urls = [...]string{
		"https://www.google.com/teapot",
		"https://example.com",
	}
	overflowMsg string
)

func getStatus(wg *sync.WaitGroup, url string, ch chan string) (err error) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
		return err
	} else {
		log.Println("Proxy: GET", url)
	}
	defer resp.Body.Close()

	ch <- resp.Status

	return nil
}

func (h *respHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	ch := make(chan string, len(urls))

	h.mu.Lock()
	if h.n == ^uint64(0) {
		overflowMsg = "more than" + " "
	} else {
		h.n++
	}
	h.mu.Unlock()

	fmt.Fprintf(w, "Backend: %s%d requests served.\n", overflowMsg, h.n)

	for _, url := range urls {
		wg.Add(1)
		go getStatus(&wg, url, ch)
	}
	wg.Wait()
	close(ch)

	for m := range ch {
		fmt.Fprintf(w, "Backend: external service responds: %q\n", m)
	}
}

func DumpReq(w http.ResponseWriter, r *http.Request) {
	resp, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Printf("DumpReq error: %s\n", err)
	}
	fmt.Fprintf(w, "Backend: header dump is:\n%s", resp)
}

func main() {
	const port string = "8080"

	log.Printf("Starting the backend service on port %s.\n", port)

	http.Handle("/", new(respHandler))
	http.HandleFunc("/dumpreq", DumpReq)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
