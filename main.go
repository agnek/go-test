package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var workers int
var endpoints string
var urlsDone int
var urlsError int

func main() {
	flag.IntVar(&workers, "workers", 5, "number of workers")
	flag.StringVar(&endpoints, "endpoints", "", "list of endpoints comma delimited")

	wg := sync.WaitGroup{}
	tasks := make(chan string)

	for i := 0; i < workers; i++ {
		go worker(wg, tasks)
	}

	for _, endpoint := range strings.Split(endpoints, ",") {
		tasks <- endpoint
	}

	wg.Wait()
	close(tasks)
	fmt.Printf("work done. successfully=%d errors=%d", urlsDone, urlsError)
}

func worker(wg sync.WaitGroup, urls chan string) error {
	wg.Add(1)
	defer wg.Done()

	for msg := range urls {
		_, err := http.Get(msg)
		if err != nil {
			urlsError += 1
			return err
		} else {
			urlsDone += 1
		}
	}

	return nil
}
