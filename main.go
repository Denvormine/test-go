package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sayHello(i)
	}
	wg.Wait()
}

func sayHello(i int) {
	fmt.Printf("hello %v\n", i)
	wg.Done()
}
