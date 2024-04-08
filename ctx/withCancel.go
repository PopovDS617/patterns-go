package ctx

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Printer(ctx context.Context, ch chan<- int) {

	for {
		rnd := rand.Int()
		time.Sleep(time.Second * 1)

		select {
		case ch <- rnd:

		case <-ctx.Done():
			close(ch)
			fmt.Println("cancelled!")
			return
		}
	}

}

func Canceller(callback func()) {

	timer := time.NewTimer(time.Second * 4)

	<-timer.C
	fmt.Println("time's up")
	callback()
}

func UseCTXWithCancel() {
	ch := make(chan int, 10)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go Canceller(cancel)
	go Printer(ctx, ch)

	for v := range ch {
		fmt.Println(v)
	}

}
