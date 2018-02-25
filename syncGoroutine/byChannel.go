package syncGoroutine

import (
	"fmt"
	"time"
)

// no buffer channel
func chanTestBasic() {
	ch := make(chan int32)
	go func() {
		fmt.Println("this is channel test")
		<-ch
	}()
	ch <- 1
}

// channel with buffer
func chanTestBuffer() {
	ch := make(chan int32, 0)

	for x := 1; x < 100000; x++ {
		go func(x int32) {
			x++
			ch <- x
		}(int32(x))
	}

	resNums := []int32{}
	for x := 1; x < 100000; x++ {
		select {
		case y := <-ch:
			resNums = append(resNums, y)
		case <-time.After(5 * time.Second):
			fmt.Println("time out")
		}
	}
	fmt.Println("results is:", len(resNums))
}

func ChanTests() {
	// chanTestBasic()
	chanTestBuffer()
}
