package ctx

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Outputter(ctx context.Context, ch chan<- string) {

	count := 1

	for {
		rnd := rand.Int()

		select {
		case ch <- fmt.Sprintf("%d - %d", count, rnd):
			fmt.Println("input #", count)
			count++
		case <-ctx.Done():
			close(ch)
			fmt.Println("time's up!")
			return
		}
	}

}

func UseCTXWithDeadline() {
	ch := make(chan string, 5)
	ctx := context.Background()
	finish := time.Now().Add(time.Second * 10)
	ctx, cancel := context.WithDeadline(ctx, finish)
	defer cancel()

	go Outputter(ctx, ch)

	for v := range ch {
		time.Sleep(time.Second * 1)
		fmt.Println("output #", strings.Split(v, " - ")[0])
	}

}
