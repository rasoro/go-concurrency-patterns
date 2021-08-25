package main

import (
	"fmt"

	waitgp "github.com/rasoro/go-concurrency-patterns/waitgroup_async_operations"
)

func main() {
	sites := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://www.google.com/search?q=golang",
	}

	resps, err := waitgp.Crawl(sites)
	if err != nil {
		panic(err)
	}

	fmt.Println("Resps received:", resps)
}
