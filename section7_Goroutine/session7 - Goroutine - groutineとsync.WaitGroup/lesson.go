// groutineとsync.WaitGroup

// goroutineは並列処理のこと

package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)
	normal("Hello")
	wg.Wait()
}
