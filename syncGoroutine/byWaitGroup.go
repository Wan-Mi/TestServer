package syncGoroutine

import (
	"fmt"
	"sync"
)

// waitGroup
func TestWaitGroup() {

	resNums := []int32{}

	var wg sync.WaitGroup
	var mt sync.Mutex
	wg.Add(100000)
	for val := 0; val < 100000; val++ {
		go func(t int32) {
			defer wg.Done()
			mt.Lock()
			resNums = append(resNums, t)
			mt.Unlock()
		}(int32(val))
	}
	wg.Wait()

	fmt.Println("results is:", len(resNums))
}
