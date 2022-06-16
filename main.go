package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

const helloCount = 20

func main() {
	for i := 0; i < helloCount; i++ {
		wg.Add(1)
		go sayHello(i)
	}
	wg.Wait()
}

func sayHello(i int) {
	fmt.Printf("hello %v\n", i)
	wg.Done()
}
