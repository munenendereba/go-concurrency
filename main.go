package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go RepeatStrings(i, &wg)
	}

	wg.Wait()
}
