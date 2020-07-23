package main

import (
	"fmt"
	"sync"
)

const None = 0
const Message = 1

type M struct {
	action int
}

type S struct {
	c chan M
	m sync.Mutex
	exit bool
}

func New() S {
	return S{
		c: make(chan M),
	}
}

func (s *S) Send(m M) {
	select {
	case s.c <- m:
		fmt.Println("sent", m)
	}
}

func (s *S) Close() {
	s.m.Lock()
	if !s.exit {
		s.exit = true
		s.c <- M{action: None}
		fmt.Println("closed")
	}
	s.m.Unlock()
}

func (s *S) Serve() {
	for range s.c {
		s.m.Lock()
		if s.exit {
			close(s.c)
		}
		s.m.Unlock()
	}
}

func C() {
	var s S = New()

	go s.Send(M{action: Message})
	go s.Close()

	s.Serve()
}
