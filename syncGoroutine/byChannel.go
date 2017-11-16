package syncGoroutine

import "fmt"
import "time"

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
	numbers := []int32{
		1, 2, 3, 4, 5,
	}
	ch := make(chan int32, len(numbers))

	for _, x := range numbers {
		go func(x int32) {
			x++
			ch <- x
		}(x)
	}

	for range numbers {
		select {
		case y := <-ch:
			fmt.Println(y)
		case <-time.After(10 * time.Second):
			fmt.Println("time out")
		}
	}
}

func ChanTests() {
	chanTestBasic()
	chanTestBuffer()
}
