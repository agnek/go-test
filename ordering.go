package main

import (
	"fmt"
	"sync/atomic"
)

func A() {
	var a, b int32

	f := func() {
		for {
			t := atomic.LoadInt32(&b)
			if t & 1 != 0 {
				continue
			}

			if atomic.CompareAndSwapInt32(&b, t, t+1) {
				a ++
				atomic.AddInt32(&b, 1)
				return
			}
		}
	}

	go f()
	go f()

	for atomic.LoadInt32(&b) < 4 {}

	fmt.Println(a)
}
