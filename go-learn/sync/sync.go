package main

import "sync"

// What this means is any goroutine calling "Inc" will acquire the lock on the Counter if they are first. The other goroutines will have to wait for it to be "Unlock"ed before getting access.

type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

// Inc the count.

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current count.

func (c *Counter) Value() int {
	return c.value
}

// If we run go vet we will still get an error. !!A Mutex must not be copied after first use.
