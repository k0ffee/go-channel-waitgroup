package main

import (
	"sync"
	"testing"
)

func TestGoogle(t *testing.T) {
	const url string = "https://www.google.com/teapot"
	const want string = "418 I'm a teapot"

	var wg sync.WaitGroup
	ch := make(chan string, 1)

	wg.Add(1)
	getStatus(&wg, url, ch)
	wg.Wait()
	close(ch)

	got := <-ch
	if got != want {
		t.Errorf("Got %q; want %q", got, want)
	}
}

func TestExample(t *testing.T) {
	const url string = "https://example.com"
	const want string = "200 OK"

	var wg sync.WaitGroup
	ch := make(chan string, 1)

	wg.Add(1)
	getStatus(&wg, url, ch)
	wg.Wait()
	close(ch)

	got := <-ch
	if got != want {
		t.Errorf("Got %q; want %q", got, want)
	}
}
