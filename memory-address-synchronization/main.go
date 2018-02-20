package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var memoryAccess sync.Mutex
	var value int

	go func() {
		memoryAccess.Lock()
		defer memoryAccess.Unlock()
		value++
	}()

	time.Sleep(1 * time.Second)
	memoryAccess.Lock()
	defer memoryAccess.Unlock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
}
