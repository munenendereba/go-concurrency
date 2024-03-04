package main

import (
	"fmt"
	"strings"
	"sync"
)

func RepeatStrings(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(strings.Repeat("*", i))
}
