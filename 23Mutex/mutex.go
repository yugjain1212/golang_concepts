package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increment() {
	p.mu.Lock()
	p.views += 1
	p.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	mypost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			mypost.increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(mypost.views)
}
