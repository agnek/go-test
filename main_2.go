package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func main2() {
	wg := sync.WaitGroup{}

	tasks := make(chan string)
	defer close(tasks)

	go func() {
		for _, url := range strings.Split(endpoints, ",") {
			tasks <- url
		}
	}()

	for i := 0; i < workers; i++ {
		go worker2(wg, tasks)
	}

	wg.Add(workers)
	wg.Wait()

	fmt.Printf("work done. successfully=%d errors=%d", done, errors)
}

func worker2(wg sync.WaitGroup, task chan string) error {
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
