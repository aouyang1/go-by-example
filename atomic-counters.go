package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	numIter := 50
	var ops uint64 = 0

	var wg sync.WaitGroup

	wg.Add(numIter)

	for i := 0; i < numIter; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 10000; j++ {
				atomic.AddUint64(&ops, 1)
			}

		}()
	}

	wg.Wait()

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
