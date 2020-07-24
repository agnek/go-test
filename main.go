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

var done int
var errors int

func main() {
	wg := sync.WaitGroup{}

	tasks := make(chan string)
	defer close(tasks)

	for i := 0; i < workers; i++ {
		go worker(wg, tasks)
	}

	for _, url := range strings.Split(endpoints, ",") {
		tasks <- url
	}

	wg.Wait()

	fmt.Printf("work done. successfully=%d errors=%d", done, errors)
}

func worker(wg sync.WaitGroup, task chan string) error {
	wg.Add(1)
	defer wg.Done()

	for url := range task {
		_, err := http.Get(url)
		if err != nil {
			errors += 1
			return err
		}

		done += 1
	}

	return nil
}

func init() {
	flag.IntVar(&workers, "workers", 5, "number of workers")
	flag.StringVar(&endpoints, "endpoints", "", "list of endpoints comma delimited")
}
