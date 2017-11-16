package syncGoroutine

import (
	"fmt"
	"sync"
)

func TestWaitGroup() {
	testSlice := []int64{
		1,
		2,
		3,
		4,
	}
	var res int64

	var wg sync.WaitGroup
	wg.Add(len(testSlice))
	for _, val := range testSlice {
		go func(t int64) {
			defer wg.Done()
			res += t

		}(val)
	}
	wg.Wait()

	fmt.Println("results is :", res)
}
