package ctx

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Logger(ctx context.Context, ch chan<- int) {

	for {
		rnd := rand.Int()

		select {
		case ch <- rnd:

		case <-ctx.Done():
			close(ch)
			fmt.Println("time's up!")
			return
		}
	}

}

func UseCTXWithTimeout() {
	ch := make(chan int, 10)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*4)
	defer cancel()

	go Logger(ctx, ch)

	for v := range ch {
		fmt.Println(len(ch))
		time.Sleep(time.Second * 1)
		fmt.Println(v)
	}

}
