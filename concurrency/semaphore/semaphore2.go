package semaphore

import (
	"fmt"
	"time"
)

func Semaphore2() {

	n := 7
	maxGoroutines := 2
	semaphore := make(chan struct{}, maxGoroutines)

	outputCh := make(chan int, n)

	for i := 0; i < n; i++ {

		go func(i int) {
			fmt.Printf("task pre semaphore %d\n", i)
			semaphore <- struct{}{}

			defer func() { <-semaphore }()

			// Simulate a task
			time.Sleep(1 * time.Second)
			outputCh <- i
			fmt.Printf("task post semaphore %d\n", i)

		}(i)
	}
	res := make([]int, 0, n)

	for i := 0; i < n; i++ {

		val := <-outputCh

		res = append(res, val)

	}
	close(outputCh)
	fmt.Println(res)
}
