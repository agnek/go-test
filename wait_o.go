package main

import (
	"fmt"
	"sync"
)

func B() {
	var wg sync.WaitGroup
	var a [2]bool

	f := func(i int) {
		a[i] = true
		wg.Done()
	}

	go func() {
		wg.Wait()
		fmt.Println(a[0], a[1])
	}()

	wg.Add(2)

	go func() {
		wg.Wait()
		fmt.Println(a[0], a[1])
	}()

	go f(0)
	go f(1)

	wg.Wait()
	fmt.Println(a[0], a[1])
}
